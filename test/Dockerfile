FROM python:3

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY test_service.py test_service.py

CMD [ "python", "./test_service.py" ]