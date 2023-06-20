# Crawler_MPSP
Project for crawlling lawsuit data avaliable in [MPSP](https://sismpconsultapublica.mpsp.mp.br/ConsultarProcedimentos/ObterProcedimentos).
Data Craw:
- cover (reference data of the lawsuit)Cancel changes
- bonds (relation to other procedures)
- doocuments (lawsuit documents)
- movements (steps of the lawsuit)
CSV file is generate with collected data.
 
## Dependencies
- [Selenium](https://github.com/tebeka/selenium#readme)
- [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/)
- [Selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/)
- [htmlquery](https://github.com/antchfx/htmlquery)

## Run
```brew install chrome driver ``` (not needed if you alredy have chrome driver)

```java -jar selenium-server-standalone.jar```

```go run main.go```

- To config a new search go in **crawler.go** file, function **DayCrawler** and alter the URL parameter on **driver.Get("")**

## Notes
- Sometimes chromedriver need a clearnce in security options on MacOS.
- Don't forget to previus install Java.
