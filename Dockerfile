FROM flokkr/base:33
ENV CONF_DIR /opt/flink/conf
ENV HADOOP_CONF_DIR /opt/flink/conf
ENV PATH $PATH:/opt/flink/bin
ADD url .
RUN wget `cat url` -O flink.tar.gz && tar zxf flink.tar.gz && rm flink.tar.gz && mv flink* flink
WORKDIR /opt/flink

