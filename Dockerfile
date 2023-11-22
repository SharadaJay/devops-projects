FROM ubuntu

RUN apt-get update && apt-get install -y openssh-server python3 sudo net-tools

RUN useradd -m -s /bin/bash -G sudo ssluser

RUN mkdir -p /home/ssluser/.ssh
COPY id_rsa.pub /home/ssluser/.ssh/authorized_keys
RUN chown -R ssluser:ssluser /home/ssluser/.ssh
RUN chmod 700 /home/ssluser/.ssh
RUN chmod 600 /home/ssluser/.ssh/authorized_keys
RUN echo 'ssluser ALL=(ALL) NOPASSWD: ALL' > /etc/sudoers.d/ssluser

ENTRYPOINT service ssh start && tail -f /dev/null