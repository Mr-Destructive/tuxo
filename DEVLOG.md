## Tuxo

## 17-03-2024

JAMStack based blog generator

Turso + Golang + HTMX + GitHub Actions


### Tech Stack

- Golang
    - Static Site Generator
    - Markdown
    - YAML / JSON / TOML

- Turso
    - LibSQL Database
    - SQLite

- HTMX
    - Component as CMS request

- GitHub
    - Actions
        - CronJob to query db every hour
        - Generate .md files from the updated posts


### Workflow for an SSG

- Load Config
- Find the md folder config
- Load templates
- Convert md to html
- Render Templates
- 

