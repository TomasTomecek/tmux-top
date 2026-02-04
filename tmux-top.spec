%global debug_package   %{nil}
%global provider        github
%global provider_tld    com
%global project         TomasTomecek
%global repo            tmux-top
# https://github.com/TomasTomecek/tmux-top
%global goipath         %{provider}.%{provider_tld}/%{project}/%{repo}
%global forgeurl        https://%{goipath}
# %%global commit          910ef1f72549a703c3c39abaefefe9a80d0b22fd
%global golicenses      LICENSE
%global godocs          README.md
# since tmux-top release archives are structured as NAME-VERSION
# 1. We need to set version
# 2. And cannot set commit macro
# 3. Place version before gometa
Version:        1.0.0
%gometa

Name:           tmux-top
Release:        %autorelease
Summary:        Monitoring information for your tmux status line.
# Automatically converted from old format: GPLv2+ - review is highly recommended.
License:        GPL-2.0-or-later
URL:            %{gourl}
# gosource macro doesn't work as it expects vTAG tagging scheme
Source0:        https://%{goipath}/archive/%{version}/%{name}-%{version}.tar.gz

BuildRequires:  make
BuildRequires:  golang(github.com/urfave/cli/v2)


%description
Monitoring information for your tmux status line.

tmux-top allows you to see your:

 * load
 * memory usage
 * network information
 * I/O

%gopkg

%prep
%goprep

%build
%gobuild -o %{gobuilddir}/bin/%{name} %{goipath}/cmd/tmux-top
# Generate manpage
%{gobuilddir}/bin/%{name} generate-man > %{name}.1


%install
%gopkginstall
install -m 0755 -vd %{buildroot}%{_bindir}
install -m 0755 -vp %{gobuilddir}/bin/%{name} %{buildroot}%{_bindir}/
# Install manpage
install -m 0755 -vd %{buildroot}%{_mandir}/man1
install -m 0644 -vp %{name}.1 %{buildroot}%{_mandir}/man1/


%check
export GOPATH=$(pwd):%{gopath}
make test


%files
%license %{golicenses}
%doc %{godocs}
%{_bindir}/%{name}
%{_mandir}/man1/%{name}.1*

%gopkgfiles


%changelog
%autochangelog
