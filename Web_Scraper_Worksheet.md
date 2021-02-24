# Web Scraper Worksheet

## Q1 Website to Scrape [1 Point]

Which website will you scrape? \
EXAMPLE: [https://news.ycombinator.com](https://dev.bukkit.org)

[My Choosen Webstie To Scrape](https://dev.bukkit.org)

## Q2 Data Elements [2 Points]

What data elements do you want? Add one or more in the space provided below. \
EXAMPLE: The title of the most popular post on the homepage.

`
Title, Image, Source Code link, Primary & Secondary Sections of the page.
`

## Q3 Selectors [2 Points]

Find and test the selectors for each data element listed above. Write each below: \
EXAMPLE: a.storylink:first-child

`#site-main > section.atf > div > div > div.project-details-container > div > div > a > img` \
`#site-main > section.atf > div > div > div.project-details-container > div > h1` \
`#site-main > section.atf > nav > div > ul > li:nth-child(6) > a` \
`#content > section > div.e-project-details-primary` \
`#content > section > div.e-project-details-secondary` \

## Q4 Review Docs Before Class  [1 Point]

1. Check out the list of features that Colly supports, as well as the best practices and examples [on the sidebar in the docs](http://go-colly.org/docs/).
2. Check out the list of features that ChromeDP supports here, as well as numerous examples in [this repo](https://github.com/chromedp/examples).
3. Which will you use? \
    Be sure to take your time to make an informed decision! Great developers allow ample time for research before choosing the framework most appropriate for the problem.
   - [ ] Colly
   - [ ] ChromeDP
   - [X] I'm not sure yet!
4. **WHY? Justify your decision in a few sentences below:**

`ChromeDP seems it would be easier to setup/configure however it seems a bit more robust than what I need. Not sure yet though...`

## Q5 Wait for Class    [1 Point]

Class will resume after an hour of working time, and then we'll take a short break.
We'll learn how to work with JSON to save the data we scrape from the web, then we'll get our web scraper projects set up properly.
If you finish early, keep working on your SSG Project. The MVP is due tonight!
