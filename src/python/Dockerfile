# Start from a small image that already has Python installed.
FROM python:3.15.2-slim

# Use /app as the working folder inside the container.
WORKDIR /app

# Copy the list of add-ons and install them.
COPY requirements.txt .
RUN pip install -r requirements.txt

# Copy the program.
COPY main.py .

# Start the program. 0.0.0.0 lets the browser on your computer reach it.
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
