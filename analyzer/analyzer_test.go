package analyzer

import "testing"

func TestLicenseType(t *testing.T) {
	tests := []struct {
		rawText string
		want    string
	}{
		{"This project is licensed under the MIT license.", "MIT"},
		{"Licensed under the Apache License, Version 2.0", "Apache"},
		{"This software is distributed under the GPL.", "GPL"},
		{"Redistribution and use in source and binary forms, with or without modification, are permitted under the BSD license.", "BSD"},
		{"Permission to use, copy, modify, and distribute this software for any purpose with or without fee is hereby granted under the ISC license.", "ISC"},
		{"This is free and unencumbered software released into the public domain under the Unlicense.", "Unlicense"},
		{"This text does not mention any license.", "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.rawText, func(t *testing.T) {
			if got := LicenseType(tt.rawText); got != tt.want {
				t.Errorf("LicenseType(%q) = %v, want %v", tt.rawText, got, tt.want)
			}
		})
	}
}
