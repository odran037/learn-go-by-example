require 'open-uri'
require 'nokogiri'
require 'enumerator'
require 'fileutils'

url = "https://gobyexample.com"
html = open(url)
doc = Nokogiri::HTML(html)
keywords = doc.xpath("//li//a/@href")
basedir = ARGV[0] || "learn-go-by-example"

Dir.mkdir basedir

keywords = keywords.enum_for(:each_with_index).map do |keyword, index|
  filename = "%02d-#{keyword}.go" % index
  path = "#{basedir}/%02d-#{keyword}" % index
  FileUtils::mkdir_p path
  File.new("#{path}/#{filename}", "w")
end
