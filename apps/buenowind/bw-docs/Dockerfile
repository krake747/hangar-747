FROM python:3.13-slim

WORKDIR /src

COPY ./requirements.txt ./mkdocs.yml ./docs /src/

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 8000

ENTRYPOINT ["mkdocs"]
CMD ["serve", "--dev-addr=0.0.0.0:8000"]
