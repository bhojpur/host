FROM registry.suse.com/suse/sle15:15.3

ENV SSL_CERT_DIR /etc/bhojpur/ssl

RUN zypper -n update && \
    zypper -n install git-core curl ca-certificates unzip mkisofs xz gzip sed tar && \
    zypper -n clean -a && \
    rm -rf /tmp/* /var/tmp/* /usr/share/doc/packages/*

RUN useradd -u 1000 bhojpur

RUN mkdir -p .bhojpur/machine/machines /etc/bhojpur/ssl /home/bhojpur && \
    chown -R bhojpur /etc/bhojpur/ssl && \
    chown -R bhojpur /home/bhojpur

COPY download_driver.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/download_driver.sh

COPY bhojpur entrypoint.sh /usr/local/bin/
RUN chmod 0777 /usr/local/bin

USER 1000
WORKDIR /home/bhojpur

ENTRYPOINT ["entrypoint.sh"]