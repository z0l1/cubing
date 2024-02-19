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
		//Right
		{
			side: Blue,
			dir:  Top,
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
	},

	//Red
	{
		//Top
		{
			side: White,
			dir:  Bottom,
		},
		//Right
		{
			side: Blue,
			dir:  Left,
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
	},

	//Blue
	{
		//Top
		{
			side: White,
			dir:  Right,
		},
		//Right
		{
			side: Orange,
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
	},

	//Orange
	{
		//Top
		{
			side: Yellow,
			dir:  Bottom,
		},
		//Right
		{
			side: Blue,
			dir:  Right,
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
	},

	//Green
	{
		//Top
		{
			side: White,
			dir:  Left,
		},
		//Right
		{
			side: Red,
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
	},

	//Yellow
	{
		//Top
		{
			side: Red,
			dir:  Bottom,
		},
		//Right
		{
			side: Blue,
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
	},
}

func getConnections(s SideColor) [4]SideConnection {
	return sideConnections[s]
}
