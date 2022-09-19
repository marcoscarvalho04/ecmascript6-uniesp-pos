# Projeto de ES - Pós frontend 

O intuíto desse projeto é apresentar as melhorias e novidades de EcmaScript para
a turma de front end do MBA Fullstack development da UNIESP. O foco aqui
não é a apresentação de HTML e CSS avançado (tanto que o leiaute não foi
bem trabalhado como deveria), mas apresentar o básico de conceitos como
promises, fetch e consumo de API's e utilização do ES para criação e manipulação
do DOM. 

# Pre-requisitos 

Para fazer a execução correta dos sistemas aqui apresentados, necessita-se: 
### 1) Docker - versão 20.10.11 ou + 
Utilizado para levantar o banco de dados postgres através do arquivo
docker-compose. 
### 2) Visual Studio Code - Recomendado
Para visualização correta do código. 
### 3) Go - versão 1.17.6 ou +
Utilizado para fazer a compilação do servidor de backend.
### 4) NPM - 8.11.0 ou + 
Utilizado para fazer a iniciação do servidor de frontend. 

# Execução 
## Servidor de frontend 
Todos os arquivos necessários para levantamento do servidor de frontend
estão sob a pasta src/. Neste caso, para inicialização do servidor faça 
a instalação do do parcel através do comando: *npm install --save-dev parcel*. 
Após isso, faça a inicialização do servidor através do comando *npm start*.

## Servidor de backend. 
Todos os arquivos necessários para levantamento do servidor de backend
estão sob a pasta server. Para inicialização do servidor, será necessário
fazer a inicialização da instância do banco postgres através do docker-compose.
Para isso, navegue até a pasta do servidor e faça o up do banco através do 
comando *docker-compose up -d*. Após isso, faça o build do servidor através do 
comando go build cmd/main/main.go ainda no nível de pasta do servidor e,
após isso, execute o main.exe gerado. 

### Observações

O foco aqui foi apenas o aprendizado de Javascript, então diversas melhorias 
foram deixadas de lado para o build no lado do servidor, além dos testes unitários. 
Para um futuro, em caso de necessidade, isso poderia ser melhorado. 




