FROM centos:7

RUN mkdir /apps
COPY ./trivybeat /apps/trivybeat

ENTRYPOINT /apps/trivybeat
CMD [ "-e", "-d", "*" ]
