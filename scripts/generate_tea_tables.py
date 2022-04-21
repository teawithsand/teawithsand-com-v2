import random


elems = {
    "tea-radius": [],
    "flow-time": [],
    "flow-time-function": [],

    "appear-time-function": [],
    "appear-time": [],
    "disappear-end-opacity": [],
    "disappear-time": [],

    "start-height": [],
    "end-height": [],

    "delay": [],

    "sand-color-start": [],
    "sand-color-end": [],

    "skew-start": [],
    "skew-end": [],
}

def trim_float(a, n=3):
    return str(round(a, n))

def random_px(a, b):
    return str(random.randint(a, b)) + "px"

def random_vh(a, b):
    return trim_float(random.uniform(a, b)) + "vh"

def random_time(a, b):
    return trim_float(random.uniform(a, b)) + "s"

def prep_float(a, unit):
    return trim_float(float(a)) + unit

def prep_int(a, unit):
    return str(int(a)) + unit

def rect_lin(a):
    if a < 0:
        return 0
    return a

def trunc_hex(a):
    return hex(a)[2:]

sand_start_colors = [
    "rgb(204, 255, 51)",
    "rgb(255, 204, 102)",
    "rgb(51, 204, 255)",
    "rgb(255, 102, 153)",
    "rgb(200, 101, 132)",
    "rgb(136, 146, 191)", # php purple
]

sand_end_colors = sand_start_colors

sand_columns = 20

def populate_table():
    for i in range(700):
        ft = random.uniform(5, 7)
        ft_avg = (5+7)/2

        elems["tea-radius"].append(prep_int(random.randint(60, 100), "px"))
        elems["flow-time"].append(prep_float(ft, "s"))
        elems["flow-time-function"].append(random.choice([
            "ease-in",
            "linear",
        ]))

        elems["appear-time-function"].append("ease-in")

        elems["appear-time"].append(prep_float(min(ft, random.uniform(.8, 1.5)), "s"))
        elems["disappear-time"].append(prep_float(min(ft, random.uniform(1, 2)), "s"))
        elems["disappear-end-opacity"].append(prep_float(random.uniform(0, 1), ""))

        elems["start-height"].append(prep_int(-100, "px"))
        elems["end-height"].append(random_vh(100, 100))
        elems["delay"].append(prep_float(random.uniform(-ft_avg, ft_avg), "s"))

        elems["sand-color-start"].append(random.choice(sand_start_colors))
        elems["sand-color-end"].append(random.choice(sand_end_colors))

        elems["skew-start"].append(prep_int(random.randint(-100, 100), "px"))
        elems["skew-end"].append(prep_int(random.randint(-20, 20), "vw"))
        

def print_table():
    print(f"$sand-columns: {sand_columns};")
    for k, v in elems.items():
        print("$" + k +":"+" "+", ".join(v) + ";")

populate_table()
print_table()