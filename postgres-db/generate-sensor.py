import csv
import numpy as np
import matplotlib.pyplot as plt

sensor_types = [
    {
        'type_abr': 'msi',
        'sensor_type': 'multi-sensor instruments',
        'basic_information': ['longitude', 'latitude', 'altitude'],
        'measurements': ['temperature','humidity','air_pressure','wind_speed']
    },
    {
        'type_abr': 'aqs',
        'sensor_type': 'air quality sensors',
        'basic_information': ['longitude', 'latitude'],
        'measurements': ['temperature','humidity','co2_level']
    },
    {
        'type_abr': 'bs',
        'sensor_type': 'basic sensor',
        'basic_information': ['longitude', 'latitude'],
        'measurements': ['temperature','humidity']
    },
]

sensor_metas = {
    'msi': {'min_distance': 20, 'n_sensor': 5},
    'aqs': {'min_distance': 10, 'n_sensor': 10},
    'bs': {'min_distance': 1, 'n_sensor': 30}
}

def haversine(longs, lats, long, lat):
    longs = np.array(longs)
    lats = np.array(lats)
    R = 6371
    phi1 = np.radians(lats)
    phi2 = np.radians(lat)
    delta_phi = np.radians(lats-lat)
    delta_lambda = np.radians(longs-long)

    a = np.sin(delta_phi/2)**2 + np.cos(phi1)*np.cos(phi2)*np.sin(delta_lambda/2)**2
    c = 2*np.arctan2(np.sqrt(a), np.sqrt(1-a))
    distance = R*c
    return distance

def is_outside(longs, lats, long, lat, min_distance):
    if len(longs)==0:
        return True
    return np.all(haversine(longs,lats,long,lat)>=min_distance)

def generate_sensor(n=100):
    rng = np.random.default_rng(0)
    candidate_lats = -6.386649 + (1-2*rng.random(1000))*20/111
    candidate_longs = 106.820701 + (1-2*rng.random(1000))*20/111
    available_ids = np.ones(1000, dtype=bool)
    ids = np.arange(1000)
    data = [['sensor_name', 'sensor_type', 'latitude', 'longitude']]
    fig, ax = plt.subplots(figsize=(8,8))

    for st in sensor_types:
        type_abr = st['type_abr']
        n_sensor = sensor_metas[type_abr]['n_sensor']
        min_distance = sensor_metas[type_abr]['min_distance']
        # print(type_abr, n_sensor, min_distance)
        # assert False
        longs = []
        lats = []
        max_retry = 5
        retry = 0
        while len(longs)<n_sensor and retry<=max_retry:
            longs = []
            lats = []
            temp_ids = available_ids.copy()
            while np.any(temp_ids) and len(longs)<n_sensor:
                id = rng.choice(ids[temp_ids])
                temp_ids[id] = False
                long = candidate_longs[id]
                lat = candidate_lats[id]
                if is_outside(longs, lats, long, lat, min_distance):
                    longs.append(long)
                    lats.append(lat)
            retry += 1
        if retry<=max_retry:
            available_ids = temp_ids.copy()
            for i in range(len(longs)):
                data.append([f"{type_abr}{i:03d}", type_abr, lats[i], longs[i]])
            ax.scatter(longs, lats, label=type_abr)
    
    ax.set_xlim(106.820701 + np.array([-1.05, 1.05])*20/111)
    ax.set_ylim(-6.386649 + np.array([-1.05, 1.05])*20/111)
    fig.legend()
    plt.show()

    with open('sensor.csv', mode='w', newline='') as file:
        writer = csv.writer(file)
        writer.writerows(data)

if __name__=="__main__":
    generate_sensor()