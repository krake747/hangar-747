FROM python:3.13-slim

WORKDIR /src

RUN pip install uv

COPY ./requirements.txt ./mkdocs.yml ./docs /src/

RUN uv pip install -r requirements.txt --system

EXPOSE 8000

ENTRYPOINT ["mkdocs"]
CMD ["serve", "--dev-addr=0.0.0.0:8000"]
