# aesir
## Under development
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)
![CSS3](https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white)
### A simple System monitor(Sysmon) EVTX inspector; search, visualize, and track Sysmon events

* **CI/CD status** ▶ [![automatic-docker-deployment](https://github.com/KnightChaser/aesir/actions/workflows/CICD.yml/badge.svg)](https://github.com/KnightChaser/aesir/actions/workflows/CICD.yml)

## Preview
You can upload your own Sysmon EVTX file with your own name
![1](./_readme_pictures/1.png)
Based on **MongoDB**, Your EVTX data will be structurally managed.
![2](./_readme_pictures/2.png)
Main page to get overall insight of your EVTX
![3](./_readme_pictures/3.png)
You can search(filter) with multiple conditions in your log file, with a simple statistics that how many results you got.
![4](./_readme_pictures/4.png)
You can search additional detailed information about the event at `Additional Information` tab that provides every detailed information(Currently focused on Sysmon(System Monitor)). Click `Go` button at the `Full metadata` column in the result table.
![5](./_readme_pictures/5.png)