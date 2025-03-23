# Science OS Demo

[![GitHub Actions CI/CD](https://github.com/plh97/SOS/actions/workflows/github-CICD-actions.yml/badge.svg)](https://github.com/plh97/SOS/actions/workflows/github-CICD-actions.yml)

<!-- ![Vercel Deploy](https://deploy-badge.vercel.app/vercel/sos-svelte) -->

## Architect

- Frontend
  - Svelte
  - TypeScript
  - Vercel
- Backend
  - Postgresql
  - Java
- Devops
  - Docker
  - Vultr
  - Terraform
  - Ansible

## Project Description

Svelte + Java project demo, once modify file, will trigger frontend auto deploy on Vercel, will trigger vultr update by github action

## Create Could Service

```bash
cd terraform/vultr
terraform init
terraform plan
terraform apply -auto-approve
terraform destroy # remove could service
```

## Init Could Service

```bash
cd ansible
ansible-playbook init-vultr.yml   # to init could env, clone project, deploy project
ansible-playbook cicd-vultr.yml   # to update project, pull latest git code, stop docker-compose, re-start docker-compose
```
