rm -rf riaas/*
swagger generate client --with-flatten=full -f swagger.yaml  -A riaas -t riaas/
for x in `find riaas -name '*.go'`
do
sed -n '
1h
1!H
$ {
   g
   s/if response.Code()\/100[^\}]*\}//g
   p
}
' $x > $x.new
mv $x.new $x
done
