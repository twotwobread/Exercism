import toml
from pathlib import Path

from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from webdriver_manager.chrome import ChromeDriverManager

CONFIG_PATH = Path(__file__).parent + "config/env.toml"
CONFIG = toml.load(CONFIG_PATH)

EMAIL = CONFIG["exercism"]["email"]
PASSWORD = CONFIG["exercism"]["password"]


def main():
    for driver in get_driver():
        get_login(driver)
        driver.implicitly_wait(15)
        hrefs = get_exercise_links(driver)
        cmds = get_commands(driver, hrefs)
        save_commands(cmds)


def get_driver():
    service = Service(ChromeDriverManager().install())
    driver = webdriver.Chrome(service=service)
    try:
        yield driver
    finally:
        driver.quit()


def get_login(driver: webdriver.Chrome):
    driver.get("https://exercism.org/users/sign_in")
    driver.implicitly_wait(10)
    button = driver.find_element(
        By.XPATH, "//form//button[@class='btn-enhanced btn-l']"
    )
    button.click()

    driver.implicitly_wait(10)
    id_box = driver.find_element(By.XPATH, "//input[@name='login']")
    id_box.send_keys(EMAIL)
    password_box = driver.find_element(By.XPATH, "//input[@name='password']")
    password_box.send_keys(PASSWORD)

    button = driver.find_element(By.XPATH, "//input[@name='commit']")
    button.click()


def get_exercise_links(driver):
    driver.get("https://exercism.org/tracks/go/exercises")
    driver.implicitly_wait(10)
    elements = driver.find_elements(By.XPATH, "//section[@class='exercises']//div/a")
    return [element.get_attribute("href") for element in elements]


def get_commands(driver, hrefs):
    cmds = []
    for href in hrefs:
        driver.get(href)
        driver.implicitly_wait(10)
        try:
            element = driver.find_element(
                By.XPATH, "//button[@class='c-copy-text-to-clipboard center-message']"
            )
        except Exception:
            continue
        clipboard_str = element.get_attribute("aria-label")
        cmds.append(clipboard_str)
    print(cmds)
    return cmds


def save_commands(cmds):
    file_path = Path(__file__).parent / "commands.txt"
    with open(file_path, "w", encoding="utf-8") as file:
        file.write("\n".join(cmds))


if __name__ == "__main__":
    main()
