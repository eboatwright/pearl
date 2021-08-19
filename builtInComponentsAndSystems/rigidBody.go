package builtInComponentsAndSystems


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
)


type RigidbodyCollisionType int
const (
	NonCollision RigidbodyCollisionType = iota
	LevelCollision
)


type Rigidbody struct {
	Velocity      pearl.Vector2
	Friction      pearl.Vector2
	Gravity       pearl.Vector2
	Grounded      float32
	CoyoteTime    float32
	CollisionType RigidbodyCollisionType
}
func (rb *Rigidbody) ID() string { return "rigidbody" }


type RigidbodySystem struct {}

func (rbs *RigidbodySystem) Update(entity *pearl.Entity, scene *pearl.Scene) {
	t := entity.GetComponent("transform").(*Transform)
	rb := entity.GetComponent("rigidbody").(*Rigidbody)
	bc := entity.GetComponent("boxCollider").(*BoxCollider)

	switch rb.CollisionType {
	case NonCollision:
		rb.Velocity.Multiply(rb.Friction)
		rb.Velocity.Add(rb.Gravity)
		t.Position.Add(pearl.Vector2Floor(rb.Velocity))
	case LevelCollision:
		if scene.FindComponent("level") == nil { return }
		level := scene.FindComponent("level").(*Level)

		tile := &pearl.Entity { ID: "tile" }
		tile.AddComponents([]pearl.IComponent {
			&Transform {},
			&BoxCollider {
				Size: pearl.Vector2 { float64(level.TileSize), float64(level.TileSize) },
			},
		})
		tileT := tile.GetComponent("transform").(*Transform)
		tileBc := tile.GetComponent("boxCollider").(*BoxCollider)

		rb.Velocity.Y *= rb.Friction.Y
		rb.Velocity.Y += rb.Gravity.Y
		t.Position.Y += float64(int64(rb.Velocity.Y))

		rb.Grounded--
		for y := 0; y < len(level.Data); y++ {
			for x := 0; x < len(level.Data[0]); x++ {
				if level.Data[y][x] > 0 {
					tileT.Position = pearl.Vector2 { float64(x * level.TileSize), float64(y * level.TileSize) }
					if EntitiesOverlap(entity, tile) {
						if rb.Velocity.Y < 0 {
							if rb.Gravity.Y < 0 {
								rb.Grounded = rb.CoyoteTime
							}
							t.Position.Y = tileT.Position.Y + tileBc.Size.Y - bc.Offset.Y
						}
						if rb.Velocity.Y > 0 {
							if rb.Gravity.Y > 0 {
								rb.Grounded = rb.CoyoteTime
							}
							t.Position.Y = tileT.Position.Y - bc.Size.Y - bc.Offset.Y
						}
						rb.Velocity.Y = 0
					}
				}
			}
		}

		rb.Velocity.X *= rb.Friction.X
		rb.Velocity.X += rb.Gravity.X
		t.Position.X += float64(int64(rb.Velocity.X))

		for y := 0; y < len(level.Data); y++ {
			for x := 0; x < len(level.Data[0]); x++ {
				if level.Data[y][x] > 0 {
					tileT.Position = pearl.Vector2 { float64(x * level.TileSize), float64(y * level.TileSize) }
					if EntitiesOverlap(entity, tile) {
						if rb.Velocity.X < 0 {
							if rb.Gravity.X < 0 {
								rb.Grounded = rb.CoyoteTime
							}
							t.Position.X = tileT.Position.X + tileBc.Size.X - bc.Offset.X
						}
						if rb.Velocity.X > 0 {
							if rb.Gravity.X > 0 {
								rb.Grounded = rb.CoyoteTime
							}
							t.Position.X = tileT.Position.X - bc.Size.X - bc.Offset.X
						}
						rb.Velocity.X = 0
					}
				}
			}
		}
	default:
		panic("invalid collision type")
	}
}

func (rbs *RigidbodySystem) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {}

func (rbs *RigidbodySystem) GetRequirements() []string {
	return []string { "transform", "rigidbody", "boxCollider" }
}