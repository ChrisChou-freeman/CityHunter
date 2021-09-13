package tool

import (
	// "fmt"
	"image"
)

func CollisionDetect(
	spritRec image.Rectangle,
	levelData *LevelData,
	vector *image.Point,
	position *image.Point,
	throwObj bool,
) bool {
	hitGround := false
	borderLessSpritRec := spritRec
	borderLessSpritRec.Min.X += 1
	borderLessSpritRec.Max.X -= 1
	borderLessSpritRec.Min.Y += 1
	borderLessSpritRec.Max.Y -= 1
	topMinTilePosition := image.Point{X: borderLessSpritRec.Min.X / TILEWIDTH, Y: borderLessSpritRec.Min.Y/TILEHEIGHT - 1}
	topMaxTilePosition := image.Point{X: borderLessSpritRec.Max.X / TILEWIDTH, Y: borderLessSpritRec.Min.Y/TILEHEIGHT - 1}
	bottomMinTilePosition := image.Point{X: borderLessSpritRec.Min.X / TILEWIDTH, Y: borderLessSpritRec.Max.Y/TILEHEIGHT + 1}
	bottomMaxTilePosition := image.Point{X: borderLessSpritRec.Max.X / TILEWIDTH, Y: borderLessSpritRec.Max.Y/TILEHEIGHT + 1}

	leftMinTilePostion := image.Point{X: borderLessSpritRec.Min.X/TILEWIDTH - 1, Y: borderLessSpritRec.Min.Y / TILEHEIGHT}
	leftMaxTilePostion := image.Point{X: borderLessSpritRec.Min.X/TILEWIDTH - 1, Y: borderLessSpritRec.Max.Y / TILEHEIGHT}
	leftBottomPosition := image.Point{X: borderLessSpritRec.Min.X/TILEWIDTH - 1, Y: borderLessSpritRec.Max.Y/TILEHEIGHT + 1}
	rightMinTilePostion := image.Point{X: borderLessSpritRec.Max.X/TILEWIDTH + 1, Y: borderLessSpritRec.Min.Y / TILEHEIGHT}
	rightMaxTilePostion := image.Point{X: borderLessSpritRec.Max.X/TILEWIDTH + 1, Y: borderLessSpritRec.Max.Y / TILEHEIGHT}
	rightBottomPosition := image.Point{X: borderLessSpritRec.Max.X/TILEWIDTH + 1, Y: borderLessSpritRec.Max.Y/TILEHEIGHT + 1}

	if bottomMinTilePosition.Y > levelData.LevelInfo["tileRowNumber"]-1 {
		return hitGround
	}

	topMinTile := -2
	topMaxTile := -2
	bottomMinTile := -2
	bottomMaxTile := -2
	leftMinTile := -2
	leftMaxTile := -2
	leftBottomTile := 2
	rightMinTile := -2
	rightMaxTile := -2
	rightBottomTile := 2

	// check left and right border
	if leftMinTilePostion.X < 0 {
		if spritRec.Min.X+vector.X < 0 {
			position.X = 0
			if throwObj {
				vector.X *= -1
			} else {
				vector.X = 0
			}
		}
	} else {
		leftMinTile = levelData.TileData[getLevalDataIndex(leftMinTilePostion, levelData)]["tile"]
		leftMaxTile = levelData.TileData[getLevalDataIndex(leftMaxTilePostion, levelData)]["tile"]
		leftBottomTile = levelData.TileData[getLevalDataIndex(leftBottomPosition, levelData)]["tile"]
	}

	if rightMinTilePostion.X > levelData.LevelInfo["tileColNumber"]-1 {
		if spritRec.Max.X+vector.X > levelData.LevelInfo["tileColNumber"]*TILEWIDTH {
			position.X = levelData.LevelInfo["tileColNumber"]*TILEWIDTH - spritRec.Dx()
			if throwObj {
				vector.X *= -1
			} else {
				vector.X = 0
			}
		}
	} else {
		rightMinTile = levelData.TileData[getLevalDataIndex(rightMinTilePostion, levelData)]["tile"]
		rightMaxTile = levelData.TileData[getLevalDataIndex(rightMaxTilePostion, levelData)]["tile"]
		rightBottomTile = levelData.TileData[getLevalDataIndex(rightBottomPosition, levelData)]["tile"]
	}

	topMinTile = levelData.TileData[getLevalDataIndex(topMinTilePosition, levelData)]["tile"]
	topMaxTile = levelData.TileData[getLevalDataIndex(topMaxTilePosition, levelData)]["tile"]
	bottomMinTile = levelData.TileData[getLevalDataIndex(bottomMinTilePosition, levelData)]["tile"]
	bottomMaxTile = levelData.TileData[getLevalDataIndex(bottomMaxTilePosition, levelData)]["tile"]

	// up
	if vector.Y < 0 {
		if SliceContainItemsOr(levelData.CollisionData["full"], []int{topMinTile, topMaxTile}) {
			if spritRec.Min.Y-1+vector.Y <= topMinTilePosition.Y*TILEHEIGHT+TILEHEIGHT {
				vector.Y *= -1
			}
		}
	}

	// down
	if vector.Y > 0 {
		if SliceContainItemsOr(levelData.CollisionData["full"], []int{bottomMinTile, bottomMaxTile}) {
			if spritRec.Max.Y+1+vector.Y >= bottomMinTilePosition.Y*TILEHEIGHT {
				position.Y = bottomMinTilePosition.Y*TILEHEIGHT - spritRec.Dy()
				vector.Y = 0
				hitGround = true
			}
		}
	}

	// fall left
	if vector.X < 0 && vector.Y > 0 {
		if SliceIndexOf(levelData.CollisionData["full"], leftBottomTile) != -1 {
			if spritRec.Min.X-1+vector.X <= leftBottomPosition.X*TILEWIDTH+TILEWIDTH &&
				spritRec.Max.Y+1+vector.Y >= leftBottomPosition.Y*TILEHEIGHT {
				position.X = leftBottomPosition.X*TILEWIDTH + TILEWIDTH
				if throwObj && vector.Y != 0 {
					vector.X *= -1
				} else {
					vector.X = 0
				}
			}
		}
	}

	//left
	if vector.X < 0 {
		if SliceContainItemsOr(levelData.CollisionData["full"], []int{leftMinTile, leftMaxTile}) {
			if spritRec.Min.X-1+vector.X <= leftMinTilePostion.X*TILEWIDTH+TILEWIDTH {
				position.X = leftMinTilePostion.X*TILEWIDTH + TILEWIDTH
				if throwObj && vector.Y != 0 {
					vector.X *= -1
				} else {
					vector.X = 0
				}
			}
		}
	}

	// fall right
	if vector.X > 0 && vector.Y > 0 {
		if SliceIndexOf(levelData.CollisionData["full"], rightBottomTile) != -1 {
			if spritRec.Max.X+1+vector.X >= rightBottomPosition.X*TILEWIDTH &&
				spritRec.Max.Y+1+vector.Y >= rightBottomPosition.Y*TILEHEIGHT {
				position.X = leftBottomPosition.X*TILEWIDTH + TILEWIDTH
				if throwObj && vector.Y != 0 {
					vector.X *= -1
				} else {
					vector.X = 0
				}
			}
		}
	}

	//right
	if vector.X > 0 {
		if SliceContainItemsOr(levelData.CollisionData["full"], []int{rightMinTile, rightMaxTile}) {
			if spritRec.Max.X+1+vector.X >= rightMinTilePostion.X*TILEWIDTH {
				position.X = rightMinTilePostion.X*TILEWIDTH - spritRec.Dx()
				if throwObj && vector.Y != 0 {
					vector.X *= -1
				} else {
					vector.X = 0
				}
			}
		}
	}

	return hitGround
}
