ARG BASE=latest
FROM flokkr/base:${BASE}
ARG URL
ENV CONF_DIR /opt/flink/conf
ENV HADOOP_CONF_DIR /opt/flink/conf
ENV PATH $PATH:/opt/flink/bin
RUN wget ${URL} -O flink.tar.gz && tar zxf flink.tar.gz && rm flink.tar.gz && mv flink* flink
WORKDIR /opt/flink

