package flatcube

import "rubik/cube"

// could probably make sides indexed so its "fancier" to get neighbours
// but this explanation took as much time as making the function below

// ...
// yeah right

// six sides 4 connections
var sideConnections = [6][4]SideConnection{
	//White
	{
		//Top
		{
			side: cube.Orange,
			dir:  Bottom,
		},
		//Right
		{
			side: cube.Blue,
			dir:  Top,
		},
		//Bottom
		{
			side: cube.Red,
			dir:  Top,
		},
		//Left
		{
			side: cube.Green,
			dir:  Top,
		},
	},

	//Red
	{
		//Top
		{
			side: cube.White,
			dir:  Bottom,
		},
		//Right
		{
			side: cube.Blue,
			dir:  Left,
		},
		//Bottom
		{
			side: cube.Yellow,
			dir:  Top,
		},
		//Left
		{
			side: cube.Green,
			dir:  Right,
		},
	},

	//Blue
	{
		//Top
		{
			side: cube.White,
			dir:  Right,
		},
		//Right
		{
			side: cube.Orange,
			dir:  Right,
		},
		//Bottom
		{
			side: cube.Yellow,
			dir:  Right,
		},
		//Left
		{
			side: cube.Red,
			dir:  Right,
		},
	},

	//Orange
	{
		//Top
		{
			side: cube.Yellow,
			dir:  Bottom,
		},
		//Right
		{
			side: cube.Blue,
			dir:  Right,
		},
		//Bottom
		{
			side: cube.White,
			dir:  Top,
		},
		//Left
		{
			side: cube.Green,
			dir:  Left,
		},
	},

	//Green
	{
		//Top
		{
			side: cube.White,
			dir:  Left,
		},
		//Right
		{
			side: cube.Red,
			dir:  Left,
		},
		//Bottom
		{
			side: cube.Yellow,
			dir:  Left,
		},
		//Left
		{
			side: cube.Orange,
			dir:  Left,
		},
	},

	//Yellow
	{
		//Top
		{
			side: cube.Red,
			dir:  Bottom,
		},
		//Right
		{
			side: cube.Blue,
			dir:  Bottom,
		},
		//Bottom
		{
			side: cube.Orange,
			dir:  Top,
		},
		//Left
		{
			side: cube.Green,
			dir:  Bottom,
		},
	},
}

func getConnections(s cube.SideColor) [4]SideConnection {
	return sideConnections[s]
}
