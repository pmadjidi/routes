package main

import "testing"

func TestAssmebleResponseFromBodies(t *testing.T) {
	body1 := `{"code":"Ok","routes":[{"legs":[{"steps":[],"summary":"","weight":263.1,"duration":260.2,"distance":1886.3}],"weight_name":"routability","weight":263.1,"duration":260.2,"distance":1886.3}],"waypoints":[{"hint":"Dv8JgCp3moUXAAAABQAAAAAAAAAgAAAAIXRPQYXNK0AAAAAAcPePQQsAAAADAAAAAAAAABAAAAA6-wAA_kvMAKlYIQM8TMwArVghAwAA7wrXLH_K","distance":4.231521214,"name":"Friedrichstraße","location":[13.388798,52.517033]},{"hint":"JEvdgVmFiocGAAAACgAAAAAAAAB3AAAAppONQOodwkAAAAAA8TeEQgYAAAAKAAAAAAAAAHcAAAA6-wAAfm7MABiJIQOCbswA_4ghAwAAXwXXLH_K","distance":2.795148358,"name":"Torstraße","location":[13.39763,52.529432]}]}`
	body2 := `{"code":"Ok","routes":[{"legs":[{"steps":[],"summary":"","weight":391,"duration":389.2,"distance":3804.3}],"weight_name":"routability","weight":391,"duration":389.2,"distance":3804.3}],"waypoints":[{"hint":"Dv8JgCp3moUXAAAABQAAAAAAAAAgAAAAIXRPQYXNK0AAAAAAcPePQQsAAAADAAAAAAAAABAAAAA6-wAA_kvMAKlYIQM8TMwArVghAwAA7wrXLH_K","distance":4.231521214,"name":"Friedrichstraße","location":[13.388798,52.517033]},{"hint":"oSkYgP___38fAAAAUQAAACYAAAAeAAAAeosKQlNOX0IQ7CZCjsMGQh8AAABRAAAAJgAAAB4AAAA6-wAASufMAOdwIQNL58wA03AhAwQAvxDXLH_K","distance":2.226580806,"name":"Platz der Vereinten Nationen","location":[13.428554,52.523239]}]}`
	bodies := make([][]byte, 0)
	bodies = append(bodies, []byte(body1))
	bodies = append(bodies, []byte(body2))
	source := "13.388860,52.517037"
	destination := []string{"13.397634,52.529407", "13.428555,52.523219"}
	resp, err := assmebleResponseFromBodies(bodies, source, destination)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if resp == nil {
		t.Fail()
	}
	dur1 := resp.Routes[0].Duration
	dur2 := resp.Routes[1].Duration
	if dur1 > dur2 {
		t.Fail()
	}
	body := bodies[0]
	bodies[0] = bodies[1]
	bodies[1] = body
	destination = []string{"13.428555,52.523219", "13.397634,52.529407"}
	resp, err = assmebleResponseFromBodies(bodies, source, destination)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if resp == nil {
		t.Fail()
	}
	dur1 = resp.Routes[0].Duration
	dur2 = resp.Routes[1].Duration
	if dur1 > dur2 {
		t.Fail()
	}
}

func TestAssmebleResponseFromBodiesMissedReq(t *testing.T) {
	body1 := `{"code":"Ok","routes":[{"legs":[{"steps":[],"summary":"","weight":263.1,"duration":260.2,"distance":1886.3}],"weight_name":"routability","weight":263.1,"duration":260.2,"distance":1886.3}],"waypoints":[{"hint":"Dv8JgCp3moUXAAAABQAAAAAAAAAgAAAAIXRPQYXNK0AAAAAAcPePQQsAAAADAAAAAAAAABAAAAA6-wAA_kvMAKlYIQM8TMwArVghAwAA7wrXLH_K","distance":4.231521214,"name":"Friedrichstraße","location":[13.388798,52.517033]},{"hint":"JEvdgVmFiocGAAAACgAAAAAAAAB3AAAAppONQOodwkAAAAAA8TeEQgYAAAAKAAAAAAAAAHcAAAA6-wAAfm7MABiJIQOCbswA_4ghAwAAXwXXLH_K","distance":2.795148358,"name":"Torstraße","location":[13.39763,52.529432]}]}`
	body2 := `{"code":"Ok","routes":[{"legs":[{"steps":[],"summary":"","weight":391,"duration":260.2,"distance":3804.3}],"weight_name":"routability","weight":391,"duration":260.2,"distance":3804.3}],"waypoints":[{"hint":"Dv8JgCp3moUXAAAABQAAAAAAAAAgAAAAIXRPQYXNK0AAAAAAcPePQQsAAAADAAAAAAAAABAAAAA6-wAA_kvMAKlYIQM8TMwArVghAwAA7wrXLH_K","distance":4.231521214,"name":"Friedrichstraße","location":[13.388798,52.517033]},{"hint":"oSkYgP___38fAAAAUQAAACYAAAAeAAAAeosKQlNOX0IQ7CZCjsMGQh8AAABRAAAAJgAAAB4AAAA6-wAASufMAOdwIQNL58wA03AhAwQAvxDXLH_K","distance":2.226580806,"name":"Platz der Vereinten Nationen","location":[13.428554,52.523239]}]}`
	bodies := make([][]byte, 0)
	bodies = append(bodies, []byte(body1))
	bodies = append(bodies, []byte(body2))
	source := "13.388860,52.517037"
	destination := []string{"13.397634,52.529407", "13.428555,52.523219"}
	resp, err := assmebleResponseFromBodies(bodies, source, destination)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	if resp == nil {
		fmt.Println("failed nil")
		t.Fail()
	}
	dur1 := resp.Routes[0].Duration
	dur2 := resp.Routes[1].Duration
	fmt.Println("dur1", dur1)
	fmt.Println("dur2", dur2)
	if dur1 != dur2 {
		t.Fail()
	}

	dist1 := resp.Routes[0].Distance
	dist2 := resp.Routes[1].Distance
	fmt.Println("dist1", dist1)
	fmt.Println("dist2", dist2)
	if dist1 > dist2 {
		t.Fail()
	}
}
