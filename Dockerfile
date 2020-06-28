ARG BASE=latest
FROM flokkr/base:${BASE}
ARG ARTIFACTDIR
ADD ${ARTIFACTDIR} /opt/flink
ENV CONF_DIR /opt/flink/conf
ENV HADOOP_CONF_DIR /opt/flink/conf
ENV PATH $PATH:/opt/flink/bin
WORKDIR /opt/flink

