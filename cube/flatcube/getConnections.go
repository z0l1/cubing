package flatcube

// could probably make sides indexed so its "fancier" to get neighbours
// but this explanation took as much time as making the function below

// ...
// yeah right

// six sides 4 connections
var sideConnections = [6][4]SideConnection{
	//WhiteSide
	{
		//Top
		{
			side: OrangeSide,
			dir:  Bottom,
		},
		//Bottom
		{
			side: RedSide,
			dir:  Top,
		},
		//Left
		{
			side: GreenSide,
			dir:  Top,
		},
		//Right
		{
			side: BlueSide,
			dir:  Top,
		},
	},

	//RedSide
	{
		//Top
		{
			side: WhiteSide,
			dir:  Bottom,
		},
		//Bottom
		{
			side: YellowSide,
			dir:  Top,
		},
		//Left
		{
			side: GreenSide,
			dir:  Right,
		},
		//Right
		{
			side: BlueSide,
			dir:  Left,
		},
	},

	//BlueSide
	{
		//Top
		{
			side: WhiteSide,
			dir:  Right,
		},
		//Bottom
		{
			side: YellowSide,
			dir:  Right,
		},
		//Left
		{
			side: RedSide,
			dir:  Right,
		},
		//Right
		{
			side: OrangeSide,
			dir:  Right,
		},
	},

	//OrangeSide
	{
		//Top
		{
			side: YellowSide,
			dir:  Bottom,
		},
		//Bottom
		{
			side: WhiteSide,
			dir:  Top,
		},
		//Left
		{
			side: GreenSide,
			dir:  Left,
		},
		//Right
		{
			side: BlueSide,
			dir:  Right,
		},
	},

	//GreenSide
	{
		//Top
		{
			side: WhiteSide,
			dir:  Left,
		},
		//Bottom
		{
			side: YellowSide,
			dir:  Left,
		},
		//Left
		{
			side: OrangeSide,
			dir:  Left,
		},
		//Right
		{
			side: RedSide,
			dir:  Left,
		},
	},

	//YellowSide
	{
		//Top
		{
			side: RedSide,
			dir:  Bottom,
		},
		//Bottom
		{
			side: OrangeSide,
			dir:  Top,
		},
		//Left
		{
			side: GreenSide,
			dir:  Bottom,
		},
		//Right
		{
			side: BlueSide,
			dir:  Bottom,
		},
	},
}

func getConnections(s Side) [4]SideConnection {
	return sideConnections[s]
}
