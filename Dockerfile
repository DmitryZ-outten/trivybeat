FROM centos:8

LABEL maitainer="Dmitry Zadorozhnyi" 

RUN mkdir /apps
COPY ./trivybeat /apps/trivybeat

WORKDIR /config
ENTRYPOINT /apps/trivybeat
CMD [ "-e", "-d", "*" ]
