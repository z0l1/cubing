package flatcube

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
			side: Orange,
			dir:  Bottom,
		},
		//Bottom
		{
			side: Red,
			dir:  Top,
		},
		//Left
		{
			side: Green,
			dir:  Top,
		},
		//Right
		{
			side: Blue,
			dir:  Top,
		},
	},

	//Red
	{
		//Top
		{
			side: White,
			dir:  Bottom,
		},
		//Bottom
		{
			side: Yellow,
			dir:  Top,
		},
		//Left
		{
			side: Green,
			dir:  Right,
		},
		//Right
		{
			side: Blue,
			dir:  Left,
		},
	},

	//Blue
	{
		//Top
		{
			side: White,
			dir:  Right,
		},
		//Bottom
		{
			side: Yellow,
			dir:  Right,
		},
		//Left
		{
			side: Red,
			dir:  Right,
		},
		//Right
		{
			side: Orange,
			dir:  Right,
		},
	},

	//Orange
	{
		//Top
		{
			side: Yellow,
			dir:  Bottom,
		},
		//Bottom
		{
			side: White,
			dir:  Top,
		},
		//Left
		{
			side: Green,
			dir:  Left,
		},
		//Right
		{
			side: Blue,
			dir:  Right,
		},
	},

	//Green
	{
		//Top
		{
			side: White,
			dir:  Left,
		},
		//Bottom
		{
			side: Yellow,
			dir:  Left,
		},
		//Left
		{
			side: Orange,
			dir:  Left,
		},
		//Right
		{
			side: Red,
			dir:  Left,
		},
	},

	//Yellow
	{
		//Top
		{
			side: Red,
			dir:  Bottom,
		},
		//Bottom
		{
			side: Orange,
			dir:  Top,
		},
		//Left
		{
			side: Green,
			dir:  Bottom,
		},
		//Right
		{
			side: Blue,
			dir:  Bottom,
		},
	},
}

func getConnections(s SideColor) [4]SideConnection {
	return sideConnections[s]
}
