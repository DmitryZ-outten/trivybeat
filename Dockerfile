FROM almalinux:9-minimal

LABEL maitainer="Dmitry Zadorozhnyi" 

RUN mkdir /apps
COPY ./trivybeat /apps/trivybeat

WORKDIR /configs
ENTRYPOINT /apps/trivybeat
CMD ["-environment" "container"]
