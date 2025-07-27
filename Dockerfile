# Use a standard base image from the Jupyter project
FROM jupyter/pyspark-notebook:spark-3.5.0

USER root

RUN mamba install -c conda-forge --yes 'pyarrow' 'boto3' && \
    mamba clean --all -f -y

ENV ICEBERG_VERSION=1.5.0
ENV HADOOP_AWS_VERSION=3.3.4

WORKDIR /opt/spark/jars
RUN wget https://repo1.maven.org/maven2/org/apache/iceberg/iceberg-spark-runtime-3.5_2.12/${ICEBERG_VERSION}/iceberg-spark-runtime-3.5_2.12-${ICEBERG_VERSION}.jar && \
    wget https://repo1.maven.org/maven2/org/apache/hadoop/hadoop-aws/${HADOOP_AWS_VERSION}/hadoop-aws-${HADOOP_AWS_VERSION}.jar

RUN chown -R ${NB_UID}:${NB_GID} /home/${NB_USER}

USER ${NB_UID}