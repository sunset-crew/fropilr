FROM rockylinux:9
COPY ./root .
RUN  /bin/bash /rocky.sh
CMD ["/usr/local/bin/supervisord","-c","/etc/supervisor/supervisord.conf"]
