要创建一个临时的 Alertmanager 容器并从中拷贝 `/etc/alertmanager` 目录，您可以使用以下步骤：

1. **启动临时 Alertmanager 容器**：
   使用 Docker 运行一个临时的 Alertmanager 容器。由于您只需要访问文件系统，您可以使用 `--entrypoint` 参数来覆盖默认的启动命令，使容器启动后保持运行而不执行 Alertmanager。

   ```bash
   docker run --name temp-alertmanager -d --entrypoint sh prom/alertmanager:v0.24.0 -c "sleep infinity"
   ```

   这会创建并启动一个名为 `temp-alertmanager` 的容器，并使其持续运行。

2. **拷贝 `/etc/alertmanager` 目录**：
   使用 Docker 的 `cp` 命令从运行中的容器拷贝 `/etc/alertmanager` 目录到您的宿主机。您可以将其拷贝到当前目录或指定的目录。

   ```bash
   docker cp temp-alertmanager:/etc/alertmanager ./alertmanager
   ```

   这会将容器内的 `/etc/alertmanager` 目录拷贝到宿主机当前目录下的 `alertmanager` 文件夹中。

3. **停止并移除临时容器**：
   一旦完成拷贝，您可以停止并移除临时容器。

   ```bash
   docker stop temp-alertmanager
   docker rm temp-alertmanager
   ```

这样，您就可以在宿主机上获得 Alertmanager 的默认配置文件，然后根据需要进行修改。如果您计划将这些配置用于之前提到的 Docker Compose 配置，确保修改后的配置文件位于 Docker Compose 文件中指定的 `./alertmanager` 路径下。