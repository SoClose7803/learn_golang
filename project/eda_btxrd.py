import os
import json
import numpy as np
import cv2
import matplotlib.pyplot as plt
import pandas as pd

# Đường dẫn đến thư mục dữ liệu
dataset_path = "/kaggle/input/btxrd-data/BTXRD/"
json_dir = os.path.join(dataset_path, "Annotations")
image_dir = os.path.join(dataset_path, "images")

# Kiểm tra số lượng tệp JSON
json_files = [f for f in os.listdir(json_dir) if f.endswith(".json")]
print(f"📊 Tổng số tệp JSON: {len(json_files)}")

# Khởi tạo danh sách lưu thông tin
annotations_data = []

# Duyệt qua từng tệp JSON
for json_file in json_files:
    json_path = os.path.join(json_dir, json_file)

    # Đọc dữ liệu từ JSON
    with open(json_path, "r") as file:
        data = json.load(file)

    # Kiểm tra nếu JSON hợp lệ
    if "imagePath" in data and "shapes" in data:
        image_path = os.path.join(image_dir, data["imagePath"])
        
        # Lưu thông tin EDA
        for shape in data["shapes"]:
            label = shape.get("label", "Unknown")
            points = np.array(shape.get("points", []), dtype=np.int32)

            # Tính diện tích bounding box (xấp xỉ)
            if len(points) > 1:
                x_min, y_min = np.min(points, axis=0)
                x_max, y_max = np.max(points, axis=0)
                area = (x_max - x_min) * (y_max - y_min)
            else:
                area = 0

            annotations_data.append([json_file, data["imagePath"], label, area])

# Chuyển danh sách thành DataFrame
df = pd.DataFrame(annotations_data, columns=["JSON File", "Image", "Label", "Area"])

# Hiển thị thống kê cơ bản
print("📌 Thống kê nhãn:")
print(df["Label"].value_counts())

print("\n📌 Thống kê diện tích vùng annotation:")
print(df["Area"].describe())

# Vẽ biểu đồ phân phối nhãn
plt.figure(figsize=(8, 4))
df["Label"].value_counts().plot(kind="bar", color="lightblue")
plt.title("Phân phối số lượng nhãn")
plt.xlabel("Loại tổn thương")
plt.ylabel("Số lượng")
plt.xticks(rotation=45)
plt.show()

# Vẽ histogram diện tích vùng annotation
plt.figure(figsize=(8, 4))
plt.hist(df["Area"], bins=20, color="salmon", edgecolor="black")
plt.title("Phân phối diện tích vùng annotation")
plt.xlabel("Diện tích (pixel)")
plt.ylabel("Số lượng")
plt.show()
