# gostitch

# Installation
```$ go get github.com/maruthibasava/gostitch```

# Usage  
```$ gostitch update``` then simply edit the `stitchconf.yml` file to your liking.

# Example

**stitchconf.yml**
```
stitch_files: 
  profile_repo: 
    extension: .sql
    directory: sql/profiles
    yield: sql/profiles
  group_repo: 
    extension: .sql
    directory: sql/groups
    yield: sql/groups
    exclude: 
      - group_create.sql
      - group_force_update.sql
```
`profile_repo` and `group_repo` are the names of the stitched files.
`yield` field is where you enter the path of the desired location of your stitched file. 

**once you are done editing the configuration file, make sure to**
```$ gostitch update```
