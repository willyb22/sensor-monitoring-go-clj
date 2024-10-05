import random
import csv

def generate_sensor(n=100):
    random.seed(0)
    filename = 'sensor.csv'
    with open(filename, mode='w', newline='') as file:
        writer = csv.writer(file)
        writer.writerow(['sensorname','latitude','longitude'])
        for i in range(n):
            sensor_name = f"sensor-{i}" 
            latitude = -6.386 + .5*(1-2*random.random())
            longitude = 106.820 + .5*(1-2*random.random())
            writer.writerow([sensor_name, latitude, longitude])
    print(f"Data has been written to {filename}")

if __name__=="__main__":
    generate_sensor()