neo4j-run:
	docker run \
		--publish=7474:7474 --publish=7687:7687 \
		--volume=${HOME}/docker/neo4j/data:/data \
		neo4j

run:
	fresh -c runner.conf
