/* Getting data from web/site we can use following methods :

1. UI Scrapping :Screen scraping is the act of copying information that shows on a digital display so it can be used for another purpose. Visual data can be collected as raw text from on-screen elements such as a text or images that appear on the desktop, in an application or on a website.
It makes a bot for this purpose.
Disadvantage: Sites can easily detect the bot and hence we can be blocked

2. RSS feed : contain data in XML format. Fastest way
An RSS (Really Simple Syndication) feed is an online file that contains details about every piece of content a site has published. Each time a site publishes a new piece of content, details about that content—including the full-text of the content or a summary, publication date, author, link, etc.

3. public API : Require Key to access api
4. Private API


GO MOD FILE :  The go. mod file is the root of dependency management in GoLang. All the modules which are needed or to be used in the project are maintained in go. mod file. For all the packages we are going to import/use in our project, it will create an entry of those modules in go.
*/


/* different packages used :
1. http : make url request
2. ioutil : reading body of the response received after making url req
3. os : immediately exit with a given status like 0,1,3
4. encoding/xml : (After getting xml formatted file/response of url request, the golang can understand in its lang only. SO it coverts into xml file in structs,etc i.e. human readable format)
unmarshaling the xml to struct defined by us
5. fmt : formatting/printing */
