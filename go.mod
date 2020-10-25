module merkle

go 1.13

require (
	merkle/merkletree v0.0.1
	merkle/socket v0.0.1
)

replace merkle/socket => ./socket

replace merkle/merkletree => ./merkletree
