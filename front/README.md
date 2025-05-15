# Starter Kit Vue.js 3 + TypeScript + TailwindCSS + ESLint + Commitlint

Este é um esqueleto de projeto configurado com Vue.js 3, TypeScript, TailwindCSS, ESLint, e Git hooks (com `simple-git-hooks` e `commitlint`). Ele visa economizar tempo ao criar novos projetos com essas tecnologias e ferramentas já configuradas.

## Tecnologias Usadas

- **Vue.js 3** - Framework para construção de interfaces de usuário.
- **TypeScript** - Superset do JavaScript para maior segurança e tipagem estática.
- **TailwindCSS** - Framework de CSS utilitário para uma estilização rápida e personalizada.
- **ESLint** - Linter para garantir a qualidade do código JavaScript/TypeScript.
- **Commitlint** - Ferramenta para garantir que as mensagens de commit sigam um padrão (Conventional Commits).
- **Husky e simple-git-hooks** - Para gerenciar hooks Git (como `pre-commit`, `commit-msg` e `pre-push`).

## Instalação

1. Clone este repositório:

   ```bash
   git clone https://github.com/seu-usuario/starter-kit-vue-ts-tailwindcss.git
   cd starter-kit-vue-ts-tailwindcss

2. Instale as dependências:
  
    ```bash
    npm install

## Git Hooks

Este projeto está configurado com **simple-git-hooks** para automatizar várias tarefas ao interagir com o repositório Git.

- **pre-commit**: Roda `lint-staged` para garantir que os arquivos modificados passem pelo ESLint antes de um commit.
- **commit-msg**: Verifica se a mensagem de commit segue o padrão do **Commitlint** (Conventional Commits).
- **pre-push**: Removido (não há script de teste configurado neste projeto).

Esses hooks garantem que o código esteja sempre limpo e que as mensagens de commit sigam um formato consistente.

---

## Scripts do `package.json`

Alguns scripts úteis estão configurados para você:

- **npm run lint**: Roda o ESLint em todo o projeto para encontrar problemas de estilo e erros de código.
- **npm run lint:fix**: Roda o ESLint e tenta corrigir automaticamente os problemas encontrados.
- **npm run dev**: Roda o servidor de desenvolvimento local (em um projeto Vue.js).
- **npm run build**: Gera uma versão otimizada para produção.
- **npm run test**: Removido (não há testes configurados).

