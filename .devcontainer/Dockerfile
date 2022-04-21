# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.209.6/containers/php/.devcontainer/base.Dockerfile

FROM ubuntu:20.04 AS nsbuild

# nsjail prelude here
RUN apt-get -y update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    autoconf \
    bison \
    flex \
    gcc \
    g++ \
    git \
    libprotobuf-dev \
    libnl-route-3-dev \
    libtool \
    make \
    pkg-config \
    protobuf-compiler \
    upx \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /
RUN git clone https://github.com/google/nsjail.git
WORKDIR /nsjail
RUN git checkout tags/3.0
RUN sed -i 's/^LDFLAGS\(.*\)/LDFLAGS\1 -static/' Makefile
RUN make && strip nsjail && upx nsjail


# FROM ubuntu:latest

# START OF GOLANG STUFF
# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
FROM golang:1.18-bullseye

# Copy library scripts to execute
COPY library-scripts/*.sh library-scripts/*.env /tmp/library-scripts/

# [Option] Install zsh
ARG INSTALL_ZSH="true"
# [Option] Upgrade OS packages to their latest versions
ARG UPGRADE_PACKAGES="true"
# Install needed packages and setup non-root user. Use a separate RUN statement to add your own dependencies.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN bash /tmp/library-scripts/common-debian.sh "${INSTALL_ZSH}" "${USERNAME}" "${USER_UID}" "${USER_GID}" "${UPGRADE_PACKAGES}" "true" "true" \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# Install Go tools
ENV GO111MODULE=auto
RUN bash /tmp/library-scripts/go-debian.sh "none" "/usr/local/go" "${GOPATH}" "${USERNAME}" "false" \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="lts/*"
ENV NVM_DIR=/usr/local/share/nvm
ENV NVM_SYMLINK_CURRENT=true \
    PATH=${NVM_DIR}/current/bin:${PATH}
RUN bash /tmp/library-scripts/node-debian.sh "${NVM_DIR}" "${NODE_VERSION}" "${USERNAME}" \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# Remove library scripts for final image
RUN rm -rf /tmp/library-scripts

# END OF GOLANG STUFF

ENV TZ=Europe/Warsaw 

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get -y install --no-install-recommends curl wget sudo ca-certificates bash unzip

RUN echo 'deb [trusted=yes] https://repo.symfony.com/apt/ /' | tee /etc/apt/sources.list.d/symfony-cli.list

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get -y install --no-install-recommends zstd imagemagick symfony-cli fish git software-properties-common

# Omited when go special stuff is used
# ARG USERNAME=vscode
# ARG USER_UID=1000
# ARG USER_GID=$USER_UID
# 
# # Create the user
# RUN groupadd --gid $USER_GID $USERNAME \
#     && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
#     #
#     # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
#     && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
#     && chmod 0440 /etc/sudoers.d/$USERNAME


# [Optional] Set the default user. Omit if you want to keep the default as root.

COPY nvm_install.sh /tmp/nvm_install.sh
RUN chmod 777 /tmp/nvm_install.sh

USER $USERNAME

RUN /tmp/nvm_install.sh

ENV TZ=Europe/Warsaw 

# ENV NVM_DIR="/home/vscode/.nvm"
# RUN [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" && \
#     nvm install lts/* && \
#     npm install -g npm && \
#     npm install -g yarn
# 
# COPY --from=nsbuild /nsjail/nsjail /usr/local/bin/nsjail

# COPY xdebug.ini /usr/local/etc/php/conf.d/xdebug.ini
# 
# RUN DEBIAN_FRONTEND=noninteractive sudo add-apt-repository ppa:ondrej/php
# 
# RUN sudo apt-get update && \
#     DEBIAN_FRONTEND=noninteractive sudo apt-get -y install --no-install-recommends \
#     php8.1-cli php8.1-apcu php8.1-xdebug php8.1-mbstring php8.1-xml php8.1-zip php8.1-imagick php8.1-gd php8.1-pgsql php8.1-intl php8.1-curl php8.1-xml \
#     python3 python3-pip
# 
# RUN mkdir /tmp/composer-install
# WORKDIR /tmp/composer-install
# 
# 
# COPY composer-setup.php composer-setup.php
# RUN php -r "if (hash_file('sha384', 'composer-setup.php') === '906a84df04cea2aa72f40b5f787e49f22d4c2f19492ac310e8cba5b96ac8b64115ac402c8cd292b8a03482574915d1a8') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;" && \
# php composer-setup.php && \
# php -r "unlink('composer-setup.php');"
# 
# RUN sudo mv composer.phar /usr/local/bin/composer
# 
# RUN rm -rf /tmp/composer-install

# RUN pip3 install docker ansible

WORKDIR /

# 
# # [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
# ARG NODE_VERSION="lts/*"
# RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi
# 
# COPY --from=nsbuild /nsjail/nsjail /usr/local/bin/nsjail
# 
# ENV TZ=Europe/Warsaw 
# 
# RUN echo 'deb [trusted=yes] https://repo.symfony.com/apt/ /' | sudo tee /etc/apt/sources.list.d/symfony-cli.list
# 
# # [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && \
#     DEBIAN_FRONTEND=noninteractive apt-get -y install --no-install-recommends zstd imagemagick symfony-cli
# 
# COPY xdebug.ini /usr/local/etc/php/conf.d/xdebug.ini

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1