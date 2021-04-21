FROM centos:8

LABEL maitainer="Dmitry Zadorozhnyi" 

RUN mkdir /apps
COPY ./trivybeat /apps/trivybeat

WORKDIR /configs
ENTRYPOINT /apps/trivybeat
CMD ["-environment" "container"]
