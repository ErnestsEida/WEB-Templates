
# Rails API

Rails api was made using build in rails command

```bash
rails new <project_name> --api --database=postgresql
```

and was adjusted accordingly so it could run on docker with postgresql database running on a seperate container.

**[!]** If you decide to not use docker, you should rewrite the database.yml file thats is in the _config_ folder. Because currently it is made that it communicates through the postgresql containers host.

