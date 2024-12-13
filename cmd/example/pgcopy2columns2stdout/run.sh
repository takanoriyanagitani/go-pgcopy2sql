#!/bin/sh

create_pgcopy(){
		echo "
				COPY (

					SELECT
					    TRUE::BOOLEAN AS active,
					    'fuji'::BYTEA AS data,
					    3.776::REAL AS height_in_km,
					    1733958275.012345::DOUBLE PRECISION AS unixtime,
					    1733958275012345::BIGINT AS unixtime_us,
					    3776::INTEGER AS height_in_meter,
						JSON_BUILD_OBJECT(
							'type', 'mount',
							'location', 'east'
						)::TEXT AS name,
					    'cafef00d-dead-beaf-face-864299792458'::UUID AS id,
						CLOCK_TIMESTAMP() AS created

					UNION ALL

					SELECT
					    FALSE::BOOLEAN AS active,
					    'takao'::BYTEA AS data,
					    0.599::REAL AS height_in_km,
					    2733958275.012345::DOUBLE PRECISION AS unixtime,
					    2733958275012345::BIGINT AS unixtime_us,
					    599::INTEGER AS height_in_meter,
						JSON_BUILD_OBJECT(
							'type', 'mount',
							'location', 'west'
						)::TEXT AS name,
					    'dafef00d-dead-beaf-face-864299792458'::UUID AS id,
						CLOCK_TIMESTAMP() AS created

					/*
					UNION ALL

					SELECT
					    NULL,
					    NULL,
					    NULL,
					    NULL,
					    NULL,
					    NULL,
					    NULL,
					    NULL,
					    NULL
					*/

				) TO STDOUT WITH BINARY
			" |
		PGUSER=$PGUSER psql \
		> ./sample.d/sample.pgcopy
}

#create_pgcopy

names=(
	active
	data
	height_in_km
	unixtime
	unixtime_us
	height_in_meter
	name
	id
	created
)

IFS=,
export ENV_KEYS_KEY="${names[*]}"
IFS=

# nullable
export active=boolean-null
export data=bytes-null
export height_in_km=float-null
export unixtime=double-null
export unixtime_us=long-null
export height_in_meter=int-null
export name=string-null
export id=uuid-null
export created=time-null

# not null
export active=boolean
export data=bytes
export height_in_km=float
export unixtime=double
export unixtime_us=long
export height_in_meter=int
export name=string
export id=uuid
export created=time

cat ./sample.d/sample.pgcopy |
	./pgcopy2columns2stdout
