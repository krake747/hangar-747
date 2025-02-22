FROM python:3.12-slim

COPY ./mkdocs.yml /src/mkdocs.yml
COPY ./requirements.txt /src/requirements.txt
COPY ./docs/ /src/docs/

WORKDIR /src

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 8000

ENTRYPOINT ["mkdocs"]
CMD ["serve", "--dev-addr=0.0.0.0:8000"]
