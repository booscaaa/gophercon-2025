names = {"Ana", "João", "Carlos", "Vinícius", "Mariana", "Lúcia"}
descriptions = {
  "Evento sensacional!",
  "Muito bom!",
  "Curti bastante a palestra.",
  "Top demais!",
  "Aprendi bastante hoje.",
  "Excelente organização!"
}

math.randomseed(os.time())

request = function()
  local name = names[math.random(#names)]
  local description = descriptions[math.random(#descriptions)]
  local body = string.format('{"name":"%s","description":"%s"}', name, description)
  wrk.method = "POST"
  wrk.headers["Content-Type"] = "application/json"
  wrk.body = body
  return wrk.format(nil)
end


-- docker run --rm -v $(pwd)/config/stress-test:/config/stress-test williamyeh/wrk \
--   -t2 -c5 -d5s --timeout 2s \
--   -s /config/stress-test/review.lua https://api-gopherconlatam.bosca.me/review