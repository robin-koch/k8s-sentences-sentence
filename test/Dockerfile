FROM python:3

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY test_name_server.py test_name_server.py

CMD [ "python", "./test_name_server.py" ]