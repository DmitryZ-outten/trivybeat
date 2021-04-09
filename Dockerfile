FROM centos:8

LABEL maitainer="Dmitriy Zadorozhnyi" 

RUN mkdir /apps
COPY ./trivybeat /apps/trivybeat

ENTRYPOINT /apps/trivybeat
CMD [ "-e", "-d", "*" ]
