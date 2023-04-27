from flask import Flask, request
import base64
import ddddocr
from PIL import Image
from io import BytesIO


app = Flask(__name__)

@app.route('/cb', methods=['POST'])
def my_post_endpoint():
    # 获取POST请求的数据
    data = request.get_data()
    print(data)
    # 从base64字符串中解码图像数据
    img_data = base64.b64decode(data)

    # 将图像数据加载到PIL Image对象中
    img = Image.open(BytesIO(img_data))
    # 将图像保存到文件中
    img.save("my_1image.png")
    ocr = ddddocr.DdddOcr(old=True)
    with open("my_1image.png", 'rb') as f:
        image = f.read()
    res = ocr.classification(image)
    return res

if __name__ == '__main__':
    app.run(port=8999, debug=True)