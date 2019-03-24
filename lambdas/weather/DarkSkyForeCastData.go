package weather

func SampleData() string {
	var data = `
{
  "latitude": 38.8697,
  "longitude": -106.9878,
  "timezone": "America/Denver",
  "currently": {
    "time": 1553453339,
    "summary": "Overcast",
    "icon": "cloudy",
    "nearestStormDistance": 5,
    "nearestStormBearing": 313,
    "precipIntensity": 0,
    "precipProbability": 0,
    "temperature": 35.91,
    "apparentTemperature": 35.91,
    "dewPoint": 23.21,
    "humidity": 0.6,
    "pressure": 1017.26,
    "windSpeed": 0.52,
    "windGust": 7.19,
    "windBearing": 306,
    "cloudCover": 0.97,
    "uvIndex": 4,
    "visibility": 7.36,
    "ozone": 311.43
  },
  "hourly": {
    "summary": "Snow (< 1 in.) this evening.",
    "icon": "snow",
    "data": [
      {
        "time": 1553450400,
        "summary": "Overcast",
        "icon": "cloudy",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 34.36,
        "apparentTemperature": 34.36,
        "dewPoint": 23.14,
        "humidity": 0.63,
        "pressure": 1017.63,
        "windSpeed": 0.84,
        "windGust": 4.72,
        "windBearing": 34,
        "cloudCover": 1,
        "uvIndex": 4,
        "visibility": 7.03,
        "ozone": 308.59
      },
      {
        "time": 1553454000,
        "summary": "Overcast",
        "icon": "cloudy",
        "precipIntensity": 0.0029,
        "precipProbability": 0.01,
        "precipAccumulation": 0.022,
        "precipType": "snow",
        "temperature": 36.25,
        "apparentTemperature": 36.25,
        "dewPoint": 23.22,
        "humidity": 0.59,
        "pressure": 1017.17,
        "windSpeed": 0.65,
        "windGust": 7.75,
        "windBearing": 289,
        "cloudCover": 0.96,
        "uvIndex": 4,
        "visibility": 7.43,
        "ozone": 312.07
      },
      {
        "time": 1553457600,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0.0012,
        "precipProbability": 0.01,
        "precipAccumulation": 0.009,
        "precipType": "snow",
        "temperature": 38.34,
        "apparentTemperature": 34.7,
        "dewPoint": 22.47,
        "humidity": 0.52,
        "pressure": 1017.24,
        "windSpeed": 4.78,
        "windGust": 10.89,
        "windBearing": 305,
        "cloudCover": 0.91,
        "uvIndex": 4,
        "visibility": 9.47,
        "ozone": 315.96
      },
      {
        "time": 1553461200,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0.0092,
        "precipProbability": 0.16,
        "precipType": "rain",
        "temperature": 39.37,
        "apparentTemperature": 34.27,
        "dewPoint": 22.21,
        "humidity": 0.5,
        "pressure": 1018.02,
        "windSpeed": 7.18,
        "windGust": 15.63,
        "windBearing": 270,
        "cloudCover": 0.73,
        "uvIndex": 4,
        "visibility": 9.97,
        "ozone": 319.06
      },
      {
        "time": 1553464800,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0.0146,
        "precipProbability": 0.23,
        "precipType": "rain",
        "temperature": 40.04,
        "apparentTemperature": 33.91,
        "dewPoint": 22.12,
        "humidity": 0.48,
        "pressure": 1018.73,
        "windSpeed": 9.5,
        "windGust": 19.14,
        "windBearing": 262,
        "cloudCover": 0.67,
        "uvIndex": 3,
        "visibility": 10,
        "ozone": 321.22
      },
      {
        "time": 1553468400,
        "summary": "Snow",
        "icon": "snow",
        "precipIntensity": 0.02,
        "precipProbability": 0.23,
        "precipAccumulation": 0.152,
        "precipType": "snow",
        "temperature": 39.58,
        "apparentTemperature": 34.23,
        "dewPoint": 22.57,
        "humidity": 0.5,
        "pressure": 1019.35,
        "windSpeed": 7.73,
        "windGust": 19.41,
        "windBearing": 297,
        "cloudCover": 0.69,
        "uvIndex": 2,
        "visibility": 6.06,
        "ozone": 322.67
      },
      {
        "time": 1553472000,
        "summary": "Possible Light Snow",
        "icon": "snow",
        "precipIntensity": 0.0184,
        "precipProbability": 0.18,
        "precipAccumulation": 0.14,
        "precipType": "snow",
        "temperature": 38.82,
        "apparentTemperature": 33.08,
        "dewPoint": 22.24,
        "humidity": 0.51,
        "pressure": 1020.11,
        "windSpeed": 8.14,
        "windGust": 19.13,
        "windBearing": 271,
        "cloudCover": 0.6,
        "uvIndex": 1,
        "visibility": 9.12,
        "ozone": 323.11
      },
      {
        "time": 1553475600,
        "summary": "Possible Light Snow",
        "icon": "snow",
        "precipIntensity": 0.0141,
        "precipProbability": 0.15,
        "precipAccumulation": 0.107,
        "precipType": "snow",
        "temperature": 37.3,
        "apparentTemperature": 32.86,
        "dewPoint": 22.07,
        "humidity": 0.54,
        "pressure": 1021.23,
        "windSpeed": 5.55,
        "windGust": 15.23,
        "windBearing": 274,
        "cloudCover": 0.61,
        "uvIndex": 0,
        "visibility": 9.15,
        "ozone": 322.35
      },
      {
        "time": 1553479200,
        "summary": "Possible Light Snow",
        "icon": "snow",
        "precipIntensity": 0.0109,
        "precipProbability": 0.12,
        "precipAccumulation": 0.083,
        "precipType": "snow",
        "temperature": 35.18,
        "apparentTemperature": 30.21,
        "dewPoint": 21.16,
        "humidity": 0.56,
        "pressure": 1022.65,
        "windSpeed": 5.74,
        "windGust": 12.65,
        "windBearing": 295,
        "cloudCover": 0.61,
        "uvIndex": 0,
        "visibility": 6.1,
        "ozone": 320.61
      },
      {
        "time": 1553482800,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.0085,
        "precipProbability": 0.09,
        "precipAccumulation": 0.065,
        "precipType": "snow",
        "temperature": 33.55,
        "apparentTemperature": 27.88,
        "dewPoint": 20.78,
        "humidity": 0.59,
        "pressure": 1024.01,
        "windSpeed": 6.26,
        "windGust": 10.89,
        "windBearing": 271,
        "cloudCover": 0.6,
        "uvIndex": 0,
        "visibility": 5.31,
        "ozone": 318.29
      },
      {
        "time": 1553486400,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.0042,
        "precipProbability": 0.05,
        "precipAccumulation": 0.032,
        "precipType": "snow",
        "temperature": 32.39,
        "apparentTemperature": 29.62,
        "dewPoint": 20.17,
        "humidity": 0.6,
        "pressure": 1025.06,
        "windSpeed": 3.1,
        "windGust": 9.07,
        "windBearing": 226,
        "cloudCover": 0.59,
        "uvIndex": 0,
        "visibility": 6.29,
        "ozone": 315.56
      },
      {
        "time": 1553490000,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.0016,
        "precipProbability": 0.03,
        "precipAccumulation": 0.013,
        "precipType": "snow",
        "temperature": 31.13,
        "apparentTemperature": 28.1,
        "dewPoint": 19.4,
        "humidity": 0.61,
        "pressure": 1026.09,
        "windSpeed": 3.17,
        "windGust": 8.15,
        "windBearing": 316,
        "cloudCover": 0.56,
        "uvIndex": 0,
        "visibility": 7.35,
        "ozone": 312.28
      },
      {
        "time": 1553493600,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.002,
        "precipProbability": 0.02,
        "precipAccumulation": 0.018,
        "precipType": "snow",
        "temperature": 29.86,
        "apparentTemperature": 25.07,
        "dewPoint": 18.58,
        "humidity": 0.62,
        "pressure": 1027.01,
        "windSpeed": 4.48,
        "windGust": 8.1,
        "windBearing": 300,
        "cloudCover": 0.53,
        "uvIndex": 0,
        "visibility": 7.82,
        "ozone": 309.02
      },
      {
        "time": 1553497200,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.0003,
        "precipProbability": 0.01,
        "precipAccumulation": 0,
        "precipType": "snow",
        "temperature": 28.77,
        "apparentTemperature": 23.39,
        "dewPoint": 17.13,
        "humidity": 0.61,
        "pressure": 1027.8,
        "windSpeed": 4.9,
        "windGust": 8.37,
        "windBearing": 291,
        "cloudCover": 0.51,
        "uvIndex": 0,
        "visibility": 8.58,
        "ozone": 305.57
      },
      {
        "time": 1553500800,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.0003,
        "precipProbability": 0.01,
        "precipAccumulation": 0,
        "precipType": "snow",
        "temperature": 27.85,
        "apparentTemperature": 21.85,
        "dewPoint": 15.86,
        "humidity": 0.6,
        "pressure": 1028.34,
        "windSpeed": 5.36,
        "windGust": 8.3,
        "windBearing": 290,
        "cloudCover": 0.51,
        "uvIndex": 0,
        "visibility": 8.62,
        "ozone": 301.98
      },
      {
        "time": 1553504400,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.0004,
        "precipProbability": 0.01,
        "precipAccumulation": 0.005,
        "precipType": "snow",
        "temperature": 27.22,
        "apparentTemperature": 23.42,
        "dewPoint": 16.55,
        "humidity": 0.64,
        "pressure": 1029.11,
        "windSpeed": 3.33,
        "windGust": 4.24,
        "windBearing": 280,
        "cloudCover": 0.72,
        "uvIndex": 0,
        "visibility": 2.98,
        "ozone": 299.24
      },
      {
        "time": 1553508000,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.0003,
        "precipProbability": 0.01,
        "precipAccumulation": 0,
        "precipType": "snow",
        "temperature": 25.95,
        "apparentTemperature": 25.95,
        "dewPoint": 15.32,
        "humidity": 0.64,
        "pressure": 1029.78,
        "windSpeed": 2.93,
        "windGust": 4.17,
        "windBearing": 291,
        "cloudCover": 0.72,
        "uvIndex": 0,
        "visibility": 3.13,
        "ozone": 297.68
      },
      {
        "time": 1553511600,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 24.46,
        "apparentTemperature": 20.44,
        "dewPoint": 14.06,
        "humidity": 0.64,
        "pressure": 1030.45,
        "windSpeed": 3.22,
        "windGust": 4.47,
        "windBearing": 283,
        "cloudCover": 0.72,
        "uvIndex": 0,
        "visibility": 3.82,
        "ozone": 296.9
      },
      {
        "time": 1553515200,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0.0002,
        "precipProbability": 0.01,
        "precipAccumulation": 0,
        "precipType": "snow",
        "temperature": 23.65,
        "apparentTemperature": 19.37,
        "dewPoint": 13.4,
        "humidity": 0.64,
        "pressure": 1030.81,
        "windSpeed": 3.32,
        "windGust": 4.85,
        "windBearing": 281,
        "cloudCover": 0.72,
        "uvIndex": 0,
        "visibility": 4.85,
        "ozone": 296.95
      },
      {
        "time": 1553518800,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 23.62,
        "apparentTemperature": 19.08,
        "dewPoint": 13.77,
        "humidity": 0.65,
        "pressure": 1030.89,
        "windSpeed": 3.5,
        "windGust": 5.03,
        "windBearing": 283,
        "cloudCover": 0.7,
        "uvIndex": 0,
        "visibility": 6.9,
        "ozone": 298.15
      },
      {
        "time": 1553522400,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 24.16,
        "apparentTemperature": 19.62,
        "dewPoint": 14.71,
        "humidity": 0.67,
        "pressure": 1030.69,
        "windSpeed": 3.55,
        "windGust": 5.3,
        "windBearing": 284,
        "cloudCover": 0.67,
        "uvIndex": 0,
        "visibility": 9.53,
        "ozone": 300.13
      },
      {
        "time": 1553526000,
        "summary": "Mostly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 25.35,
        "apparentTemperature": 20.69,
        "dewPoint": 15.19,
        "humidity": 0.65,
        "pressure": 1030.12,
        "windSpeed": 3.77,
        "windGust": 6.17,
        "windBearing": 278,
        "cloudCover": 0.62,
        "uvIndex": 1,
        "visibility": 10,
        "ozone": 301.88
      },
      {
        "time": 1553529600,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 28.58,
        "apparentTemperature": 28.58,
        "dewPoint": 14.71,
        "humidity": 0.56,
        "pressure": 1028.82,
        "windSpeed": 1.32,
        "windGust": 8.33,
        "windBearing": 237,
        "cloudCover": 0.5,
        "uvIndex": 3,
        "visibility": 10,
        "ozone": 303.15
      },
      {
        "time": 1553533200,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 32.03,
        "apparentTemperature": 27.52,
        "dewPoint": 13.77,
        "humidity": 0.46,
        "pressure": 1027.14,
        "windSpeed": 4.58,
        "windGust": 11.05,
        "windBearing": 327,
        "cloudCover": 0.33,
        "uvIndex": 5,
        "visibility": 10,
        "ozone": 304.25
      },
      {
        "time": 1553536800,
        "summary": "Clear",
        "icon": "clear-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 35.67,
        "apparentTemperature": 29.89,
        "dewPoint": 13.13,
        "humidity": 0.39,
        "pressure": 1025.71,
        "windSpeed": 7.04,
        "windGust": 12.93,
        "windBearing": 283,
        "cloudCover": 0.19,
        "uvIndex": 6,
        "visibility": 10,
        "ozone": 304.48
      },
      {
        "time": 1553540400,
        "summary": "Clear",
        "icon": "clear-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 38.25,
        "apparentTemperature": 32.72,
        "dewPoint": 13.01,
        "humidity": 0.35,
        "pressure": 1024.73,
        "windSpeed": 7.51,
        "windGust": 13.44,
        "windBearing": 275,
        "cloudCover": 0.13,
        "uvIndex": 8,
        "visibility": 10,
        "ozone": 303.02
      },
      {
        "time": 1553544000,
        "summary": "Clear",
        "icon": "clear-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 40.99,
        "apparentTemperature": 36,
        "dewPoint": 13.21,
        "humidity": 0.32,
        "pressure": 1024.03,
        "windSpeed": 7.62,
        "windGust": 13.11,
        "windBearing": 275,
        "cloudCover": 0.05,
        "uvIndex": 8,
        "visibility": 10,
        "ozone": 300.79
      },
      {
        "time": 1553547600,
        "summary": "Clear",
        "icon": "clear-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 43.25,
        "apparentTemperature": 38.85,
        "dewPoint": 13.64,
        "humidity": 0.3,
        "pressure": 1023.5,
        "windSpeed": 7.4,
        "windGust": 12.27,
        "windBearing": 285,
        "cloudCover": 0,
        "uvIndex": 6,
        "visibility": 10,
        "ozone": 299.42
      },
      {
        "time": 1553551200,
        "summary": "Clear",
        "icon": "clear-day",
        "precipIntensity": 0.0002,
        "precipProbability": 0.03,
        "precipType": "rain",
        "temperature": 45.14,
        "apparentTemperature": 41.97,
        "dewPoint": 14.44,
        "humidity": 0.29,
        "pressure": 1023.03,
        "windSpeed": 5.85,
        "windGust": 10.69,
        "windBearing": 307,
        "cloudCover": 0,
        "uvIndex": 4,
        "visibility": 10,
        "ozone": 299.69
      },
      {
        "time": 1553554800,
        "summary": "Clear",
        "icon": "clear-day",
        "precipIntensity": 0.0003,
        "precipProbability": 0.03,
        "precipType": "rain",
        "temperature": 45.91,
        "apparentTemperature": 43.3,
        "dewPoint": 15.46,
        "humidity": 0.29,
        "pressure": 1022.77,
        "windSpeed": 5.17,
        "windGust": 8.55,
        "windBearing": 250,
        "cloudCover": 0,
        "uvIndex": 2,
        "visibility": 10,
        "ozone": 300.83
      },
      {
        "time": 1553558400,
        "summary": "Clear",
        "icon": "clear-day",
        "precipIntensity": 0.0003,
        "precipProbability": 0.02,
        "precipType": "rain",
        "temperature": 44.95,
        "apparentTemperature": 42.69,
        "dewPoint": 15.95,
        "humidity": 0.31,
        "pressure": 1023.02,
        "windSpeed": 4.44,
        "windGust": 6.62,
        "windBearing": 274,
        "cloudCover": 0.01,
        "uvIndex": 1,
        "visibility": 10,
        "ozone": 301.92
      },
      {
        "time": 1553562000,
        "summary": "Clear",
        "icon": "clear-day",
        "precipIntensity": 0.0003,
        "precipProbability": 0.02,
        "precipType": "rain",
        "temperature": 40.83,
        "apparentTemperature": 40.83,
        "dewPoint": 15.67,
        "humidity": 0.36,
        "pressure": 1024.09,
        "windSpeed": 1.74,
        "windGust": 5.04,
        "windBearing": 293,
        "cloudCover": 0.06,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 303.01
      },
      {
        "time": 1553565600,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 36.16,
        "apparentTemperature": 36.16,
        "dewPoint": 14.94,
        "humidity": 0.41,
        "pressure": 1025.66,
        "windSpeed": 2.49,
        "windGust": 3.67,
        "windBearing": 98,
        "cloudCover": 0.14,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 304.06
      },
      {
        "time": 1553569200,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 32.32,
        "apparentTemperature": 32.32,
        "dewPoint": 13.91,
        "humidity": 0.46,
        "pressure": 1026.95,
        "windSpeed": 1.09,
        "windGust": 2.98,
        "windBearing": 97,
        "cloudCover": 0.18,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 304.34
      },
      {
        "time": 1553572800,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 30.63,
        "apparentTemperature": 30.63,
        "dewPoint": 12.8,
        "humidity": 0.47,
        "pressure": 1027.73,
        "windSpeed": 1.12,
        "windGust": 3.42,
        "windBearing": 108,
        "cloudCover": 0.14,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 303.25
      },
      {
        "time": 1553576400,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 29.76,
        "apparentTemperature": 29.76,
        "dewPoint": 11.46,
        "humidity": 0.46,
        "pressure": 1028.25,
        "windSpeed": 2.47,
        "windGust": 4.5,
        "windBearing": 112,
        "cloudCover": 0.07,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 301.39
      },
      {
        "time": 1553580000,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 28.96,
        "apparentTemperature": 24.98,
        "dewPoint": 9.88,
        "humidity": 0.44,
        "pressure": 1028.53,
        "windSpeed": 3.66,
        "windGust": 5.34,
        "windBearing": 114,
        "cloudCover": 0.03,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 299.68
      },
      {
        "time": 1553583600,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 28.35,
        "apparentTemperature": 23.89,
        "dewPoint": 7.64,
        "humidity": 0.41,
        "pressure": 1028.58,
        "windSpeed": 3.98,
        "windGust": 5.61,
        "windBearing": 116,
        "cloudCover": 0.06,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 298.52
      },
      {
        "time": 1553587200,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 27.83,
        "apparentTemperature": 23.3,
        "dewPoint": 5.12,
        "humidity": 0.37,
        "pressure": 1028.42,
        "windSpeed": 3.97,
        "windGust": 5.63,
        "windBearing": 118,
        "cloudCover": 0.1,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 297.65
      },
      {
        "time": 1553590800,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 27.59,
        "apparentTemperature": 23.12,
        "dewPoint": 3.2,
        "humidity": 0.34,
        "pressure": 1028.2,
        "windSpeed": 3.89,
        "windGust": 5.58,
        "windBearing": 119,
        "cloudCover": 0.14,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 296.98
      },
      {
        "time": 1553594400,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 26.5,
        "apparentTemperature": 22.07,
        "dewPoint": 2.04,
        "humidity": 0.34,
        "pressure": 1028,
        "windSpeed": 3.73,
        "windGust": 5.43,
        "windBearing": 120,
        "cloudCover": 0.16,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 296.42
      },
      {
        "time": 1553598000,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 25.13,
        "apparentTemperature": 20.83,
        "dewPoint": 1.52,
        "humidity": 0.35,
        "pressure": 1027.73,
        "windSpeed": 3.48,
        "windGust": 5.18,
        "windBearing": 119,
        "cloudCover": 0.18,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 296.18
      },
      {
        "time": 1553601600,
        "summary": "Clear",
        "icon": "clear-night",
        "precipIntensity": 0.0002,
        "precipProbability": 0.01,
        "precipAccumulation": 0,
        "precipType": "snow",
        "temperature": 24.67,
        "apparentTemperature": 20.58,
        "dewPoint": 2.29,
        "humidity": 0.37,
        "pressure": 1027.25,
        "windSpeed": 3.29,
        "windGust": 5.03,
        "windBearing": 120,
        "cloudCover": 0.2,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 296.34
      },
      {
        "time": 1553605200,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-night",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 25.83,
        "apparentTemperature": 22.12,
        "dewPoint": 5.15,
        "humidity": 0.41,
        "pressure": 1026.53,
        "windSpeed": 3.14,
        "windGust": 4.74,
        "windBearing": 125,
        "cloudCover": 0.25,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 297.4
      },
      {
        "time": 1553608800,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 27.73,
        "apparentTemperature": 24.49,
        "dewPoint": 9.28,
        "humidity": 0.45,
        "pressure": 1025.58,
        "windSpeed": 3,
        "windGust": 4.52,
        "windBearing": 132,
        "cloudCover": 0.31,
        "uvIndex": 0,
        "visibility": 10,
        "ozone": 299.13
      },
      {
        "time": 1553612400,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 30.64,
        "apparentTemperature": 27.55,
        "dewPoint": 12.86,
        "humidity": 0.47,
        "pressure": 1024.5,
        "windSpeed": 3.16,
        "windGust": 5.18,
        "windBearing": 142,
        "cloudCover": 0.35,
        "uvIndex": 1,
        "visibility": 10,
        "ozone": 300.07
      },
      {
        "time": 1553616000,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 34.86,
        "apparentTemperature": 31.43,
        "dewPoint": 15.29,
        "humidity": 0.44,
        "pressure": 1023.11,
        "windSpeed": 3.95,
        "windGust": 7.4,
        "windBearing": 160,
        "cloudCover": 0.35,
        "uvIndex": 3,
        "visibility": 10,
        "ozone": 299.55
      },
      {
        "time": 1553619600,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 39.71,
        "apparentTemperature": 36,
        "dewPoint": 17.25,
        "humidity": 0.4,
        "pressure": 1021.56,
        "windSpeed": 5.17,
        "windGust": 10.47,
        "windBearing": 180,
        "cloudCover": 0.34,
        "uvIndex": 5,
        "visibility": 10,
        "ozone": 298.38
      },
      {
        "time": 1553623200,
        "summary": "Partly Cloudy",
        "icon": "partly-cloudy-day",
        "precipIntensity": 0,
        "precipProbability": 0,
        "temperature": 44.49,
        "apparentTemperature": 40.91,
        "dewPoint": 18.87,
        "humidity": 0.35,
        "pressure": 1020.09,
        "windSpeed": 6.32,
        "windGust": 13.09,
        "windBearing": 195,
        "cloudCover": 0.33,
        "uvIndex": 6,
        "visibility": 10,
        "ozone": 297.92
      }
    ]
  },
  "daily": {
    "summary": "Snow (2–5 in.) today and Friday, with high temperatures peaking at 51°F on Wednesday.",
    "icon": "snow",
    "data": [
      {
        "time": 1553407200,
        "summary": "Snow (< 1 in.) in the evening.",
        "icon": "snow",
        "sunriseTime": 1553432774,
        "sunsetTime": 1553477073,
        "moonPhase": 0.64,
        "precipIntensity": 0.0051,
        "precipIntensityMax": 0.02,
        "precipIntensityMaxTime": 1553468400,
        "precipProbability": 0.5,
        "precipAccumulation": 0.802,
        "precipType": "snow",
        "temperatureHigh": 40.04,
        "temperatureHighTime": 1553464800,
        "temperatureLow": 23.62,
        "temperatureLowTime": 1553518800,
        "apparentTemperatureHigh": 36.25,
        "apparentTemperatureHighTime": 1553454000,
        "apparentTemperatureLow": 19.08,
        "apparentTemperatureLowTime": 1553518800,
        "dewPoint": 20.68,
        "humidity": 0.64,
        "pressure": 1018.7,
        "windSpeed": 2.12,
        "windGust": 19.41,
        "windGustTime": 1553468400,
        "windBearing": 270,
        "cloudCover": 0.72,
        "uvIndex": 4,
        "uvIndexTime": 1553446800,
        "visibility": 7.66,
        "ozone": 310.34,
        "temperatureMin": 24.75,
        "temperatureMinTime": 1553432400,
        "temperatureMax": 40.04,
        "temperatureMaxTime": 1553464800,
        "apparentTemperatureMin": 24.75,
        "apparentTemperatureMinTime": 1553432400,
        "apparentTemperatureMax": 36.25,
        "apparentTemperatureMaxTime": 1553454000
      },
      {
        "time": 1553493600,
        "summary": "Mostly cloudy in the morning.",
        "icon": "partly-cloudy-day",
        "sunriseTime": 1553519078,
        "sunsetTime": 1553563531,
        "moonPhase": 0.67,
        "precipIntensity": 0.0002,
        "precipIntensityMax": 0.002,
        "precipIntensityMaxTime": 1553493600,
        "precipProbability": 0.05,
        "precipAccumulation": 0.044,
        "precipType": "snow",
        "temperatureHigh": 45.91,
        "temperatureHighTime": 1553554800,
        "temperatureLow": 24.67,
        "temperatureLowTime": 1553601600,
        "apparentTemperatureHigh": 43.3,
        "apparentTemperatureHighTime": 1553554800,
        "apparentTemperatureLow": 20.58,
        "apparentTemperatureLowTime": 1553601600,
        "dewPoint": 14.61,
        "humidity": 0.49,
        "pressure": 1027.1,
        "windSpeed": 3.35,
        "windGust": 13.44,
        "windGustTime": 1553540400,
        "windBearing": 284,
        "cloudCover": 0.34,
        "uvIndex": 8,
        "uvIndexTime": 1553540400,
        "visibility": 8.59,
        "ozone": 301.71,
        "temperatureMin": 23.62,
        "temperatureMinTime": 1553518800,
        "temperatureMax": 45.91,
        "temperatureMaxTime": 1553554800,
        "apparentTemperatureMin": 19.08,
        "apparentTemperatureMinTime": 1553518800,
        "apparentTemperatureMax": 43.3,
        "apparentTemperatureMaxTime": 1553554800
      },
      {
        "time": 1553580000,
        "summary": "Partly cloudy throughout the day.",
        "icon": "partly-cloudy-day",
        "sunriseTime": 1553605383,
        "sunsetTime": 1553649988,
        "moonPhase": 0.7,
        "precipIntensity": 0.0002,
        "precipIntensityMax": 0.0003,
        "precipIntensityMaxTime": 1553655600,
        "precipProbability": 0.06,
        "precipType": "rain",
        "temperatureHigh": 50.55,
        "temperatureHighTime": 1553637600,
        "temperatureLow": 35.89,
        "temperatureLowTime": 1553688000,
        "apparentTemperatureHigh": 50.55,
        "apparentTemperatureHighTime": 1553637600,
        "apparentTemperatureLow": 32.02,
        "apparentTemperatureLowTime": 1553688000,
        "dewPoint": 15.07,
        "humidity": 0.4,
        "pressure": 1021.86,
        "windSpeed": 3.91,
        "windGust": 17.25,
        "windGustTime": 1553634000,
        "windBearing": 178,
        "cloudCover": 0.35,
        "uvIndex": 7,
        "uvIndexTime": 1553626800,
        "visibility": 10,
        "ozone": 305.61,
        "temperatureMin": 24.67,
        "temperatureMinTime": 1553601600,
        "temperatureMax": 50.55,
        "temperatureMaxTime": 1553637600,
        "apparentTemperatureMin": 20.58,
        "apparentTemperatureMinTime": 1553601600,
        "apparentTemperatureMax": 50.55,
        "apparentTemperatureMaxTime": 1553637600
      },
      {
        "time": 1553666400,
        "summary": "Mostly cloudy throughout the day.",
        "icon": "partly-cloudy-day",
        "sunriseTime": 1553691688,
        "sunsetTime": 1553736445,
        "moonPhase": 0.74,
        "precipIntensity": 0.0004,
        "precipIntensityMax": 0.004,
        "precipIntensityMaxTime": 1553731200,
        "precipProbability": 0.06,
        "precipType": "rain",
        "temperatureHigh": 51.01,
        "temperatureHighTime": 1553720400,
        "temperatureLow": 34.92,
        "temperatureLowTime": 1553774400,
        "apparentTemperatureHigh": 51.01,
        "apparentTemperatureHighTime": 1553720400,
        "apparentTemperatureLow": 33.3,
        "apparentTemperatureLowTime": 1553781600,
        "dewPoint": 21.03,
        "humidity": 0.43,
        "pressure": 1014.57,
        "windSpeed": 5.92,
        "windGust": 27.07,
        "windGustTime": 1553720400,
        "windBearing": 219,
        "cloudCover": 0.61,
        "uvIndex": 6,
        "uvIndexTime": 1553709600,
        "visibility": 10,
        "ozone": 308.55,
        "temperatureMin": 35.89,
        "temperatureMinTime": 1553688000,
        "temperatureMax": 51.01,
        "temperatureMaxTime": 1553720400,
        "apparentTemperatureMin": 32.02,
        "apparentTemperatureMinTime": 1553688000,
        "apparentTemperatureMax": 51.01,
        "apparentTemperatureMaxTime": 1553720400
      },
      {
        "time": 1553752800,
        "summary": "Mostly cloudy throughout the day.",
        "icon": "partly-cloudy-day",
        "sunriseTime": 1553777993,
        "sunsetTime": 1553822902,
        "moonPhase": 0.77,
        "precipIntensity": 0.0004,
        "precipIntensityMax": 0.0018,
        "precipIntensityMaxTime": 1553835600,
        "precipProbability": 0.04,
        "precipType": "rain",
        "temperatureHigh": 50.59,
        "temperatureHighTime": 1553806800,
        "temperatureLow": 32.06,
        "temperatureLowTime": 1553868000,
        "apparentTemperatureHigh": 50.59,
        "apparentTemperatureHighTime": 1553806800,
        "apparentTemperatureLow": 28.98,
        "apparentTemperatureLowTime": 1553868000,
        "dewPoint": 16.14,
        "humidity": 0.36,
        "pressure": 1013.41,
        "windSpeed": 5.06,
        "windGust": 30.92,
        "windGustTime": 1553806800,
        "windBearing": 233,
        "cloudCover": 0.5,
        "uvIndex": 6,
        "uvIndexTime": 1553796000,
        "visibility": 9.21,
        "ozone": 309.79,
        "temperatureMin": 34.92,
        "temperatureMinTime": 1553774400,
        "temperatureMax": 50.59,
        "temperatureMaxTime": 1553806800,
        "apparentTemperatureMin": 33.3,
        "apparentTemperatureMinTime": 1553781600,
        "apparentTemperatureMax": 50.59,
        "apparentTemperatureMaxTime": 1553806800
      },
      {
        "time": 1553839200,
        "summary": "Snow (1–3 in.) throughout the day.",
        "icon": "snow",
        "sunriseTime": 1553864299,
        "sunsetTime": 1553909359,
        "moonPhase": 0.8,
        "precipIntensity": 0.0133,
        "precipIntensityMax": 0.0219,
        "precipIntensityMaxTime": 1553911200,
        "precipProbability": 0.58,
        "precipAccumulation": 2.701,
        "precipType": "snow",
        "temperatureHigh": 33.12,
        "temperatureHighTime": 1553864400,
        "temperatureLow": 21.53,
        "temperatureLowTime": 1553954400,
        "apparentTemperatureHigh": 31.85,
        "apparentTemperatureHighTime": 1553907600,
        "apparentTemperatureLow": 15.92,
        "apparentTemperatureLowTime": 1553954400,
        "dewPoint": 15.52,
        "humidity": 0.51,
        "pressure": 1015.45,
        "windSpeed": 2.04,
        "windGust": 10.09,
        "windGustTime": 1553889600,
        "windBearing": 260,
        "cloudCover": 0.94,
        "uvIndex": 4,
        "uvIndexTime": 1553886000,
        "visibility": 1.1,
        "ozone": 381.75,
        "temperatureMin": 26.9,
        "temperatureMinTime": 1553922000,
        "temperatureMax": 36.99,
        "temperatureMaxTime": 1553839200,
        "apparentTemperatureMin": 22.22,
        "apparentTemperatureMinTime": 1553886000,
        "apparentTemperatureMax": 36.99,
        "apparentTemperatureMaxTime": 1553839200
      },
      {
        "time": 1553925600,
        "summary": "Foggy starting in the evening.",
        "icon": "fog",
        "sunriseTime": 1553950604,
        "sunsetTime": 1553995817,
        "moonPhase": 0.83,
        "precipIntensity": 0.0048,
        "precipIntensityMax": 0.0081,
        "precipIntensityMaxTime": 1553994000,
        "precipProbability": 0.21,
        "precipAccumulation": 1.226,
        "precipType": "snow",
        "temperatureHigh": 35.33,
        "temperatureHighTime": 1553986800,
        "temperatureLow": 13.77,
        "temperatureLowTime": 1554033600,
        "apparentTemperatureHigh": 35.33,
        "apparentTemperatureHighTime": 1553986800,
        "apparentTemperatureLow": 11.71,
        "apparentTemperatureLowTime": 1554040800,
        "dewPoint": 8.34,
        "humidity": 0.45,
        "pressure": 1024.75,
        "windSpeed": 3.25,
        "windGust": 7.77,
        "windGustTime": 1553968800,
        "windBearing": 61,
        "cloudCover": 0.93,
        "uvIndex": 3,
        "uvIndexTime": 1553965200,
        "visibility": 4.99,
        "ozone": 408.17,
        "temperatureMin": 21.53,
        "temperatureMinTime": 1553954400,
        "temperatureMax": 35.33,
        "temperatureMaxTime": 1553986800,
        "apparentTemperatureMin": 15.92,
        "apparentTemperatureMinTime": 1553954400,
        "apparentTemperatureMax": 35.33,
        "apparentTemperatureMaxTime": 1553986800
      },
      {
        "time": 1554012000,
        "summary": "Partly cloudy in the morning.",
        "icon": "partly-cloudy-night",
        "sunriseTime": 1554036910,
        "sunsetTime": 1554082274,
        "moonPhase": 0.86,
        "precipIntensity": 0.0003,
        "precipIntensityMax": 0.0038,
        "precipIntensityMaxTime": 1554012000,
        "precipProbability": 0.05,
        "precipAccumulation": 0.119,
        "precipType": "snow",
        "temperatureHigh": 46.45,
        "temperatureHighTime": 1554073200,
        "temperatureLow": 19.27,
        "temperatureLowTime": 1554120000,
        "apparentTemperatureHigh": 43.9,
        "apparentTemperatureHighTime": 1554073200,
        "apparentTemperatureLow": 19.27,
        "apparentTemperatureLowTime": 1554120000,
        "dewPoint": 8.72,
        "humidity": 0.45,
        "pressure": 1028.33,
        "windSpeed": 3.09,
        "windGust": 9.8,
        "windGustTime": 1554073200,
        "windBearing": 345,
        "cloudCover": 0.21,
        "uvIndex": 8,
        "uvIndexTime": 1554058800,
        "visibility": 8.67,
        "ozone": 328.01,
        "temperatureMin": 13.77,
        "temperatureMinTime": 1554033600,
        "temperatureMax": 46.45,
        "temperatureMaxTime": 1554073200,
        "apparentTemperatureMin": 11.71,
        "apparentTemperatureMinTime": 1554040800,
        "apparentTemperatureMax": 43.9,
        "apparentTemperatureMaxTime": 1554073200
      }
    ]
  },
  "flags": {
    "sources": [
      "nearest-precip",
      "nwspa",
      "cmc",
      "gfs",
      "hrrr",
      "icon",
      "isd",
      "madis",
      "nam",
      "sref",
      "darksky"
    ],
    "nearest-station": 0.849,
    "darksky-unavailable": "Dark Sky covers the given location, but all stations are currently unavailable.",
    "units": "us"
  },
  "offset": -6
}
`
	return data
}
