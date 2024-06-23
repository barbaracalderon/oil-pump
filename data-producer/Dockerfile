FROM python:3.10-slim as python-base

ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1 \
    PIP_NO_CACHE_DIR=off \
    PIP_DISABLE_PIP_VERSION_CHECK=on \
    PIP_DEFAULT_TIMEOUT=100

RUN rm -rf /code

WORKDIR /code

FROM python-base as builder-base

RUN apt-get update \
    && apt-get install --no-install-recommends -y \
    curl \
    build-essential

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

WORKDIR /code

EXPOSE 8000

# Command to run the application
CMD ["python", "src/main.py"]
