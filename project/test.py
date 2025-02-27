#đọc file .json (Chiến)
import os
import json
import numpy as np
import cv2
import matplotlib.pyplot as plt

# Đường dẫn đến tệp JSON
json_path = "/kaggle/input/btxrd-data/BTXRD/Annotations/IMG000006.json"

# Kiểm tra sự tồn tại của tệp JSON
if os.path.exists(json_path):
    with open(json_path, "r") as file:
        data = json.load(file)

    # Kiểm tra khóa "imagePath" trong JSON
    image_file = data.get("imagePath")
    if image_file:
        image_path = os.path.join("/kaggle/input/btxrd-data/BTXRD/images", image_file)

        # Kiểm tra sự tồn tại của tệp ảnh
        if os.path.exists(image_path):
            image = cv2.imread(image_path)

            if image is not None:
                # Kiểm tra khóa "shapes" trong JSON
                shapes = data.get("shapes", [])
                if shapes:
                    # Màu hồng (RGB: 255, 105, 180) -> BGR: (180, 105, 255) trong OpenCV
                    color_pink = (180, 105, 255)

                    # Vẽ annotation trên ảnh
                    for shape in shapes:
                        label = shape.get("label", "Unknown")  # Nhãn vùng chú thích
                        points = np.array(shape.get("points", []), dtype=np.int32)  # Tọa độ vùng

                        if len(points) == 0:
                            print(f"Lỗi: Không có điểm nào cho nhãn {label}")
                            continue  # Bỏ qua nếu không có điểm

                        # Vẽ vùng annotation với màu hồng
                        cv2.polylines(image, [points], isClosed=True, color=color_pink, thickness=2)

                        # Hiển thị nhãn trên ảnh
                        x, y = points[0]
                        cv2.putText(image, label, (x, y - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.6, color_pink, 2)

                    # Hiển thị ảnh đã annotate
                    plt.figure(figsize=(8, 8))
                    plt.imshow(cv2.cvtColor(image, cv2.COLOR_BGR2RGB))
                    plt.axis("off")
                    plt.show()
                else:
                    print("Lỗi: Không tìm thấy 'shapes' trong JSON")
            else:
                print("Lỗi: OpenCV không thể đọc ảnh")
        else:
            print(f"Lỗi: Không tìm thấy ảnh {image_path}")
    else:
        print("Lỗi: Không tìm thấy khóa 'imagePath' trong JSON")
else:
    print(f"Lỗi: Không tìm thấy tệp JSON {json_path}")

