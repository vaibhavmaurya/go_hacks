from datetime import datetime, timedelta
from airflow import DAG
from airflow.operators.python import PythonOperator
import time
import boto3
import json
from datetime import datetime

# AWS credentials - Use environment variables or a secure method to set these
aws_access_key_id = ''
aws_secret_access_key = ''
aws_region = 'ap-south-1'

session = boto3.session.Session(
        aws_access_key_id=aws_access_key_id,
        aws_secret_access_key=aws_secret_access_key,
        region_name=aws_region
    )

client = session.client('events')

# def create_boto3_session(aws_access_key_id, aws_secret_access_key, aws_region):
#     return boto3.session.Session(
#         aws_access_key_id=aws_access_key_id,
#         aws_secret_access_key=aws_secret_access_key,
#         region_name=aws_region
#     )

def build_custom_context_dict(context):
    # Initialize the custom context dictionary
    custom_context_dict = {
        "web_host": context.get('conf', {}).get('webserver', 'web_server_host'),
        "web_port": context.get('conf', {}).get('webserver', 'web_server_port'),
        "Origin": "pipeline_dag",
        "dag": {
            "dag_id": context.get('dag', {}).dag_id if 'dag' in context else None,
            "start_date": str(context.get('dag', {}).start_date) if 'dag' in context else None,
            "end_date": str(context.get('dag', {}).end_date) if 'dag' in context else None,
        },
        "dag_run": {
            "run_id": context.get('dag_run', {}).run_id if 'dag_run' in context else None,
            "state": context.get('dag_run', {}).state if 'dag_run' in context else None,
            "execution_date": str(context.get('dag_run', {}).execution_date) if 'dag_run' in context else None,
            "start_date": str(context.get('dag_run', {}).start_date) if 'dag_run' in context else None,
            "end_date": str(context.get('dag_run', {}).end_date) if 'dag_run' in context else None,
        },
        "execution_date": str(context.get('execution_date')) if 'execution_date' in context else None,
        "logical_date": str(context.get('logical_date')) if 'logical_date' in context else None,
        "next_execution_date": str(context.get('next_execution_date')) if 'next_execution_date' in context else None,
        "run_id": context.get('run_id'),
        "task": {
            "task_id": context.get('task', {}).task_id if 'task' in context else None,
            "owner": context.get('task', {}).owner if 'task' in context else None,
            "start_date": str(context.get('task', {}).start_date) if 'task' in context else None,
            "end_date": str(context.get('task', {}).end_date) if 'task' in context else None,
            "retries": context.get('task', {}).retries if 'task' in context else None,
        },
        "ti": {
            "task_id": context.get('ti', {}).task_id if 'ti' in context else None,
            "execution_date": str(context.get('ti', {}).execution_date) if 'ti' in context else None,
            "start_date": str(context.get('ti', {}).start_date) if 'ti' in context else None,
            "end_date": str(context.get('ti', {}).end_date) if 'ti' in context else None,
            "state": context.get('ti', {}).state if 'ti' in context else None,
            "try_number": context.get('ti', {}).try_number if 'ti' in context else None,
        },
    }

    return custom_context_dict


# Usage within a callback function:
# custom_context_dict = build_custom_context_dict(context)
# You can then use this dictionary as needed, e.g., log it, send it via an API call, etc.



def send_event_to_eventbridge(detail_type, detail):
    # client = session.client('events')
    response = client.put_events(
        Entries=[
            {
                'Time': datetime.now(),
                'Source': 'airflow.dag',
                'Resources': [
                    'airflow_dag',
                ],
                'DetailType': detail_type,
                'Detail': json.dumps(build_custom_context_dict(detail)),
                'EventBusName': 'default'
            },
        ]
    )
    print("-----------SENDING EVENT-----------")
    print(response)
    return response

def handle_callback(context):
    # session = create_boto3_session(aws_access_key_id, aws_secret_access_key, aws_region)
    # print('Calling Callback')
    # print(context)
    send_event_to_eventbridge('Airflow DAG Satus', context)


def sleep_for_seconds(seconds, **kwargs):
    time.sleep(seconds)


def start_dag(**kwargs):
    print("DAG has started. Initialization task is complete.")

def end_dag(**kwargs):
    print("All tasks are done. DAG has completed.")

default_args = {
    'owner': 'airflow',
    'depends_on_past': False,
    'email_on_failure': False,
    'email_on_retry': False,
    'retries': 1,
    'retry_delay': timedelta(minutes=1),
    # 'on_success_callback': handle_callback,
    # 'on_failure_callback': handle_callback,
    # 'sla_miss_callback' : handle_callback,
    # 'on_retry_callback' : handle_callback,
    'on_execute_callback' : handle_callback
}

with DAG(
    'parallel_and_sequential_tasks',
    default_args=default_args,
    description='A DAG with two parallel tasks and one sequential task, including sleep of different times',
    schedule_interval=None,  # This schedules the DAG to run every 5 minutes
    start_date=datetime(2024, 2, 1),
    catchup=False,
    on_success_callback=handle_callback,
    on_failure_callback=handle_callback,
) as dag:
    

    # start_task = PythonOperator(
    #     task_id='start_task',
    #     python_callable=start_dag,
    # )

    # Parallel tasks
    sleep_5_seconds = PythonOperator(
        task_id='sleep_5_seconds',
        python_callable=sleep_for_seconds,
        op_kwargs={'seconds': 5},
    )
    
    sleep_10_seconds = PythonOperator(
        task_id='sleep_10_seconds',
        python_callable=sleep_for_seconds,
        op_kwargs={'seconds': 10},
    )
    
    # Sequential task that follows the parallel tasks
    sleep_15_seconds = PythonOperator(
        task_id='sleep_15_seconds',
        python_callable=sleep_for_seconds,
        op_kwargs={'seconds': 15},
        trigger_rule='all_done',  # Ensures this runs after all parallel tasks are done
    )

    # end_task_success = PythonOperator(
    #     task_id='end_task_success',
    #     python_callable=end_dag,
    #     trigger_rule='all_success',  # This ensures it only runs if all tasks are successful
    # )

    # end_task_failed = PythonOperator(
    #     task_id='end_task_failed',
    #     python_callable=end_dag,
    #     trigger_rule='one_failed',  # This ensures it only runs if any tasks fails
    # )
    
    start_task >> [sleep_5_seconds, sleep_10_seconds] >> sleep_15_seconds
    # [sleep_5_seconds, sleep_10_seconds, sleep_15_seconds] >> end_task_failed
