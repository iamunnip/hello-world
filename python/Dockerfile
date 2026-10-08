FROM python:3.14-slim
WORKDIR /app
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY src/main.py ./
EXPOSE 8000
CMD ["flask", "--app", "main", "run", "--host", "0.0.0.0", "--port", "8000"]
