import json
import sys
import time

from playwright.sync_api import Playwright, expect, sync_playwright


def run(playwright: Playwright, account, pwd) -> None:
    browser = playwright.chromium.launch(headless=False)
    context = browser.new_context()
    page = context.new_page()
    page.goto("https://business.oceanengine.com/login?appKey=51")
    # 加载状态
    page.wait_for_load_state()
    page.locator("use").nth(1).click()
    page.get_by_placeholder("请输入邮箱").fill(account)
    page.get_by_placeholder("密码").fill(pwd)
    page.get_by_role("button", name="登录").click()
    page.wait_for_load_state()
    #  获得登录后 cookie
    time.sleep(3)
    cookieStr = ""
    cookies = context.cookies()
    for v in cookies:
        cookieStr += v.get("name") + "=" + v.get("value") + ";"
        # print(v.get("name"), "=", v.get("value"), ";")
    page.goto("https://ad.oceanengine.com/pages/promotion.html?aadvid=1759708720305223#/ad")
    time.sleep(3)
    page.wait_for_load_state()
    # storage = context.storage_state()
    # with open("cookie.json", "w") as f:
    #     f.write(json.dumps(storage))
    cookies = context.cookies()
    for v in cookies:
        cookieStr += v.get("name") + "=" + v.get("value") + ";"
    print(cookieStr)
    context.close()
    browser.close()


with sync_playwright() as playwright:
    # 763557856@qq.com
    # !qaz*963Qcrh
    # print("参数个数为:", len(sys.argv), "个参数。")
    # print("参数列表:", str(sys.argv))
    # print(sys.argv[1], sys.argv[2])
    run(playwright, sys.argv[1], sys.argv[2])
