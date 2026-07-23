from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

import requests
import time

API_URL = "http://localhost:8080/shorts"
options = webdriver.ChromeOptions()
driver = webdriver.Chrome(options=options)
driver.get("https://www.youtube.com/shorts")

wait = WebDriverWait(driver, 20)
body = wait.until(
    EC.presence_of_element_located((By.TAG_NAME, "body"))
)

history = set()

while True:

    time.sleep(10)

    try:
        title = driver.find_element(
            By.CSS_SELECTOR,
            "h2.ytShortsVideoTitleViewModelShortsVideoTitle"
        ).text

        current_url = driver.current_url

        if current_url not in history:
            history.add(current_url)
            payload = {
                "title": title,
                "url": current_url
            }

            requests.post(API_URL, json=payload, timeout=10)
            print(payload)

    except Exception:
        print("title not found")

    body.send_keys(Keys.ARROW_DOWN)

# while True:

#     time.sleep(10)

#     try :
#         title = driver.find_element(By.CSS_SELECTOR,"h2.ytShortsVideoTitleViewModelShortsVideoTitle").text
#         current_url = driver.current_url

#         if current_url not in history:
#             history.add(current_url)
#             payload = {
#                 "title": title,
#                 "url": current_url
#             }
#     except :
#         print("title not found")
#         continue
#         body.send_keys(Keys.ARROW_DOWN)
        
#     try:
#         requests.post(API_URL, json=payload, timeout=10)
#         print(payload)
#     except Exception as e:
#         print(e)

#     body.send_keys(Keys.ARROW_DOWN)