import pytest, requests

BASE = "http://localhost:8080"

def test_list_grades_ok():
    r = requests.get(f"{BASE}/grades")
    assert r.status_code == 200
    assert isinstance(r.json(), list)

def test_create_get_update_delete_grade():
    payload = {"student_id": 1, "course": "English", "score": 81}
    r = requests.post(f"{BASE}/grades", json=payload)
    assert r.status_code == 201
    gid = r.json()["id"]

    r = requests.get(f"{BASE}/grades/{gid}")
    assert r.status_code == 200

    r = requests.put(f"{BASE}/grades/{gid}", json={"student_id":1,"course":"English","score":85})
    assert r.status_code == 200

    r = requests.delete(f"{BASE}/grades/{gid}")
    assert r.status_code == 204
